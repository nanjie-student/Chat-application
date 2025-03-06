package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "chat-app/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your username: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user)

	go getMessageHistory(client, user)

	for {
		fmt.Print("Chat Type - (1) Private (2) Group (exit to quit): ")
		chatType, _ := reader.ReadString('\n')
		chatType = strings.TrimSpace(chatType)

		if chatType == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if chatType == "1" {
			fmt.Print("Enter recipient: ")
			receiver, _ := reader.ReadString('\n')
			receiver = strings.TrimSpace(receiver)

			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			sendMessage(client, user, receiver, message)
		} else if chatType == "2" {
			fmt.Print("Enter group name: ")
			group, _ := reader.ReadString('\n')
			group = strings.TrimSpace(group)

			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			sendGroupMessage(client, user, group, message)
		}
	}
}

// 发送消息
func sendMessage(client pb.ChatServiceClient, sender, receiver, content string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.MessageRequest{
		Sender:   sender,
		Receiver: receiver,
		Content:  content,
	}

	res, err := client.SendMessage(ctx, req)
	if err != nil {
		log.Fatalf("SendMessage failed: %v", err)
	}

	fmt.Printf(" Message Sent! Status: %s, Content: %s\n", res.Status, res.Content)
}
func getMessageHistory(client pb.ChatServiceClient, user string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	req := &pb.HistoryRequest{User: user}
	stream, err := client.GetMessageHistory(ctx, req)
	if err != nil {
		log.Fatalf("GetMessageHistory failed: %v", err)
	}

	fmt.Printf("📜 Chat history for %s:\n", user)
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("💬 %s\n", msg.Content)
	}
}

func sendGroupMessage(client pb.ChatServiceClient, sender, group, content string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GroupMessageRequest{
		Sender:    sender,
		GroupName: group,
		Content:   content,
	}

	client.SendGroupMessage(ctx, req)
}
