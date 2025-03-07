# Chat-application
Developed a real-time chat app using microservices architecture, featuring user authentication, real-time messaging, and message history.

1.gRPC 服务器
定义 proto 文件，包括 SendMessage、GetMessageHistory、SendGroupMessage、GetGroupMessages 方法。
生成 chat.pb.go 和 chat_grpc.pb.go，并在 server.go 里实现 gRPC 服务逻辑。
运行 go run main.go，确保 gRPC 服务器正确监听 50051 端口。
运行 grpcurl -plaintext localhost:50051 list，确认 proto.ChatService 已注册。
尝试 grpcurl 发送消息，发现 gRPC 没有返回数据（仍需排查）。

2.WebSocket 服务器
实现 WebSocket 服务器，监听 8080/ws，允许前端建立 WebSocket 连接。
客户端可以发送消息到 WebSocket 服务器，并转发给所有连接的客户端。
尝试用 wscat -c ws://localhost:8080/ws 测试 WebSocket，检查是否能正常发送和接收消息。

3.HTTP 服务器
实现 ServerFrontend()，托管 React 静态文件。
在 go run main.go 终端，看到 HTTP server started on port 8080，说明 HTTP 服务器已运行。
尝试访问 http://localhost:8080，确认前端页面是否正确加载。

4.目前存在的问题
    4.1gRPC SendMessage 方法没有返回数据
    grpcurl 发送请求后没有响应，需要检查 server.go 里 SendMessage 是否正确 return &msg, nil。
    检查 go run main.go 终端是否打印 "Private message from Alice to Bob: Hello, Bob!"，如果没有，说明 gRPC 请求没有正确处理。
    4.2WebSocket 需要进一步测试
    需要 wscat 测试是否能够接收 gRPC 服务器的消息推送。