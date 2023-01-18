package main

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/productFeed", websocket.Handler(server.handleWSProductFeed))

	if err := http.ListenAndServe(":8050", nil); err != nil {
		log.Fatal(err)
	}
}
