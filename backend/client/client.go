package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "Chat-application/proto"

	"google.golang.org/grpc"
)

func main() {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	// 用户输入用户名
	fmt.Print("Enter your username: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user)

	// 获取私聊和群聊消息
	go getMessageHistory(client, user)

	for {
		fmt.Print("Chat Type - (1) Private (2) Group (3) Get Group History (exit to quit): ")
		chatType, _ := reader.ReadString('\n')
		chatType = strings.TrimSpace(chatType)

		if chatType == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if chatType == "1" { // 私聊
			fmt.Print("Enter recipient: ")
			receiver, _ := reader.ReadString('\n')
			receiver = strings.TrimSpace(receiver)

			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			sendMessage(client, user, receiver, message)
		} else if chatType == "2" { // 发送群聊消息
			fmt.Print("Enter group name: ")
			group, _ := reader.ReadString('\n')
			group = strings.TrimSpace(group)

			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)

			sendGroupMessage(client, user, group, message)
		} else if chatType == "3" { // 获取群聊记录
			fmt.Print("Enter group name: ")
			group, _ := reader.ReadString('\n')
			group = strings.TrimSpace(group)

			getGroupMessages(client, group)
		}
	}
}

// 发送私聊消息
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

	fmt.Printf("Message Sent! Status: %s, Content: %s\n", res.Status, res.Content)
}

// 获取私聊历史
func getMessageHistory(client pb.ChatServiceClient, user string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	req := &pb.HistoryRequest{User: user}
	stream, err := client.GetMessageHistory(ctx, req)
	if err != nil {
		log.Fatalf("GetMessageHistory failed: %v", err)
	}

	fmt.Printf("Chat history for %s:\n", user)
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("%s\n", msg.Content)
	}
}

// 发送群聊消息
func sendGroupMessage(client pb.ChatServiceClient, sender, group, content string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GroupMessageRequest{
		Sender:    sender,
		GroupName: group,
		Content:   content,
	}

	res, err := client.SendGroupMessage(ctx, req)
	if err != nil {
		log.Fatalf("SendGroupMessage failed: %v", err)
	}

	fmt.Printf("Group Message Sent! Status: %s, Content: %s\n", res.Status, res.Content)
}

// 获取群聊历史
func getGroupMessages(client pb.ChatServiceClient, group string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	req := &pb.GroupHistoryRequest{GroupName: group}
	stream, err := client.GetGroupMessages(ctx, req)
	if err != nil {
		log.Fatalf("GetGroupMessages failed: %v", err)
	}

	fmt.Printf("Chat history for group %s:\n", group)
	for {
		msg, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("%s\n", msg.Content)
	}
}
