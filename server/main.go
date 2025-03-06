package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "chat-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
	messages      []pb.MessageResponse
	groupMessages map[string][]pb.MessageResponse
	mu            sync.Mutex
}

func (s *ChatService) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	msg := pb.MessageResponse{
		Status:  "Delivered",
		Content: fmt.Sprintf("[Private] %s → %s: %s", req.Sender, req.Receiver, req.Content),
	}

	s.messages = append(s.messages, msg)
	log.Printf("Private message from %s to %s: %s", req.Sender, req.Receiver, req.Content)
	return &msg, nil
}

func (s *ChatService) GetMessageHistory(req *pb.HistoryRequest, stream pb.ChatService_GetMessageHistoryServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, msg := range s.messages {
		if err := stream.Send(&msg); err != nil {
			return err
		}
	}

	log.Printf("Sent private chat history to %s", req.User)
	return nil
}

func (s *ChatService) SendGroupMessage(ctx context.Context, req *pb.GroupMessageRequest) (*pb.MessageResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	msg := pb.MessageResponse{
		Status:  "Delivered",
		Content: fmt.Sprintf("[Group: %s] %s: %s", req.GroupName, req.Sender, req.Content),
	}

	s.groupMessages[req.GroupName] = append(s.groupMessages[req.GroupName], msg)
	log.Printf("Group message in %s from %s: %s", req.GroupName, req.Sender, req.Content)

	return &msg, nil
}

func (s *ChatService) GetGroupMessages(req *pb.GroupHistoryRequest, stream pb.ChatService_GetGroupMessagesServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, msg := range s.groupMessages[req.GroupName] {
		if err := stream.Send(&msg); err != nil {
			return err
		}
	}

	log.Printf("Sent chat history for group %s", req.GroupName)
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chatService := &ChatService{
		groupMessages: make(map[string][]pb.MessageResponse),
	}

	pb.RegisterChatServiceServer(grpcServer, chatService)
	reflection.Register(grpcServer)

	log.Println("gRPC server started on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
