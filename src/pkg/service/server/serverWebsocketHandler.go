package server

import (
	"fmt"
	"golang.org/x/net/websocket"
)

func (s Server) HandleWSAccount(conn *websocket.Conn) {
	fmt.Println("new incoming connection from client:", conn.RemoteAddr())
	s.Conns[conn] = true

	s.readLoopAccount(conn)
}

func (s Server) HandleWSChat(conn *websocket.Conn) {
	fmt.Println("new incoming connection from client:", conn.RemoteAddr())
	s.Conns[conn] = true

	s.readLoopChat(conn)
}
