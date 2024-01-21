package server

import "golang.org/x/net/websocket"

type Server struct {
	Conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		Conns: make(map[*websocket.Conn]bool),
	}
}