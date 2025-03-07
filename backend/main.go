package main

import (
	"log"
	"net/http"

	"Chat-application/server"
)

func main() {
	//启动 gRPC 服务器（异步）
	go server.StartGRPCServer()
	//启动 WebSocket 服务器（异步）
	go server.StartWebSocketServer()
	//托管前端
	server.ServerFrontend()

	log.Println("HTTP server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
