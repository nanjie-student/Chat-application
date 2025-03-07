package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// ✅ WebSocket 连接管理
var clients = make(map[*websocket.Conn]bool) // 客户端连接列表
var broadcast chan string                    // 消息广播
var mu sync.Mutex                            // 互斥锁，防止并发问题

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// 让 gRPC 服务器向 WebSocket 客户端发送消息
var SendToWebSocket = func(message string) {
	broadcast <- message
}

// 启动 WebSocket 服务器
func StartWebSocketServer() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Println("WebSocket server started on port 8080")
}

// 处理 WebSocket 连接
func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	mu.Lock()
	clients[ws] = true
	mu.Unlock()

	for {
		var msg struct {
			Content string `json:"content"`
		}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket read error:", err)

			mu.Lock()
			delete(clients, ws)
			mu.Unlock()

			break
		}
		log.Printf("Received WebSocket message: %s", msg.Content)
		broadcast <- msg.Content
	}
}

// 发送 WebSocket 消息
func handleMessages() {
	for {
		msg := <-broadcast

		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("WebSocket write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
