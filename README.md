## WebSocket Chat

This repository contains a simple chat application that uses WebSockets to communicate between the client and the server.

## Prerequisites
- Go
- golang.org/x/net/websocket
- net/http

## Getting Started

1. Clone the repository
```bash
git clone https://github.com/yourusername/golang-websocket-chat.git
``


let socket = new WebSocket("ws://localhost:8050/ws")
socket.onmessage = event => {console.log("received from server: ", event.data)}
socket.send("Hello from brave tab")
let socket = new WebSocket("ws://localhost:8050/productFeed")