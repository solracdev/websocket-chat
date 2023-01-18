package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"sync"
	"time"
)

type Server struct {
	mutex      sync.RWMutex
	connexions map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		connexions: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connexion from client: ", ws.RemoteAddr())
	s.mutex.RLock()
	defer s.mutex.Unlock()
	s.connexions[ws] = true
	s.readConn(ws)
}

func (s *Server) handleWSProductFeed(ws *websocket.Conn) {
	fmt.Println("new incoming connexion from client to product feed: ", ws.RemoteAddr())
	ticker := time.NewTicker(5 * time.Second)
	for {
		payload := fmt.Sprintf("product feed data -> %d \n", time.Now().UnixNano())
		_, err := ws.Write([]byte(payload))
		if err != nil {
			fmt.Println("product feed error:", err)
			return
		}
		<-ticker.C
	}
}

func (s *Server) readConn(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}

		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.connexions {
		go func(conn *websocket.Conn) {
			if _, err := conn.Write(b); err != nil {
				fmt.Println("write error:", err)
			}
		}(ws)
	}
}
