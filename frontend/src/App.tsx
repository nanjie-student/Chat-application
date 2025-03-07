import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';
import React, { useState, useEffect } from "react";
import socket from "./ws";

const App: React.FC = () => {
  const [message, setMessage] = useState<string>("");
  const [messages, setMessages] = useState<string[]>([]);

  // 当 WebSocket 收到消息时，更新聊天记录
  useEffect(() => {
    socket.onmessage = (event: MessageEvent) => {
      setMessages((prev: string[]) => [...prev, event.data]);
    };
  }, []);

  // 发送消息
  const sendMessage = () => {
    if (message.trim()) {
      const data = JSON.stringify({ content: message }); // ✅ 发送 JSON
      socket.send(data);
      setMessage("");
    }
  };

  return (
    <div style={{ textAlign: "center", padding: "20px" }}>
      <h1>WebSock Chat</h1>
      <div style={{ minHeight: "200px", border: "1px solid #ccc", padding: "10px" }}>
        {messages.map((msg, index) => (
          <p key={index}>{msg}</p>
        ))}
      </div>
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        placeholder="Type a message..."
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
};

export default App;