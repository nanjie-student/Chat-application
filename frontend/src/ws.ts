const socket = new WebSocket("ws://localhost:8080/ws");

// 监听 WebSocket 连接成功
socket.onopen = () => {
  console.log("Connected to WebSocket server");
};

// 监听 WebSocket 接收消息
socket.onmessage = (event: MessageEvent) => {
  console.log("New WebSocket message:", event.data);
};
// 监听 WebSocket 发生错误
socket.onerror = (error: Event) => {
  console.error("WebSocket error:", error);
};

// 监听 WebSocket 断开连接
socket.onclose = () => {
    console.log("🔌 WebSocket connection closed");
};

export default socket;
