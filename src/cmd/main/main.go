package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

/* Broadcast message to all connected clients */
func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error:", err)
			}
		}(ws)
	}
}

func main() {
	signupEndpoint()
	structureEndpoint()
}

func signupEndpoint() {
	server := NewServer()
	http.Handle("/account", websocket.Handler(server.handleWS))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting signup server:", err)
	}
}

func structureEndpoint() {
	server := NewServer()
	http.Handle("/chats", websocket.Handler(server.handleWS))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting structure server:", err)
	}
}
