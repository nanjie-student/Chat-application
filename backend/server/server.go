package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "Chat-application/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	//让 Go 识别 WebSocket 代码
)

//让 Go 识别前端代码

// ChatService 实现 gRPC 服务器
type ChatService struct {
	pb.UnimplementedChatServiceServer
	messages      []pb.MessageResponse
	groupMessages map[string][]pb.MessageResponse
	mu            sync.Mutex
}

// 托管前端 React 静态文件
// func serveFrontend() {
// 	fs := http.FileServer(http.Dir("./frontend/build"))
// 	http.Handle("/", fs)
// }

// gRPC 方法：发送消息
func (s *ChatService) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	s.mu.Lock()
	//defer s.mu.Unlock()

	msg := pb.MessageResponse{
		Status:  "Delivered",
		Content: fmt.Sprintf("[Private] %s → %s: %s", req.Sender, req.Receiver, req.Content),
	}

	s.messages = append(s.messages, msg)
	s.mu.Unlock() //立即释放锁，避免影响 WebSocket 处理

	// **通知 WebSocket 客户端**
	//发送 WebSocket 通知（异步执行，防止阻塞 gRPC）
	go func(content string) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered from WebSocket panic:", r)
			}
		}()
		if SendToWebSocket != nil {
			log.Println("Sending message to WebSocket clients:", content)
			SendToWebSocket(content)
		}
	}(msg.Content)

	// **通知 WebSocket 客户端有新消息**
	//broadcast <- msg.Content

	log.Printf("Private message from %s to %s: %s", req.Sender, req.Receiver, req.Content)
	return &msg, nil
}

// gRPC 方法：获取私聊历史
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

// gRPC 方法：发送群聊消息
func (s *ChatService) SendGroupMessage(ctx context.Context, req *pb.GroupMessageRequest) (*pb.MessageResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	msg := pb.MessageResponse{
		Status:  "Delivered",
		Content: fmt.Sprintf("[Group: %s] %s: %s", req.GroupName, req.Sender, req.Content),
	}

	s.groupMessages[req.GroupName] = append(s.groupMessages[req.GroupName], msg)

	// **通知 WebSocket 客户端有新群聊消息**

	// **通知 WebSocket 客户端**
	if SendToWebSocket != nil {
		SendToWebSocket(msg.Content)
	}

	log.Printf("Group message in %s from %s: %s", req.GroupName, req.Sender, req.Content)
	return &msg, nil
}

// gRPC 方法：获取群聊历史
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

// 启动服务器
func StartGRPCServer() {
	//fmt.Println("Starting gRPC server on port 50051...") // 确保日志输出

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	chatService := &ChatService{
		groupMessages: make(map[string][]pb.MessageResponse),
	}

	pb.RegisterChatServiceServer(grpcServer, chatService)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server started on port 50051") // 确保这个日志会打印

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
