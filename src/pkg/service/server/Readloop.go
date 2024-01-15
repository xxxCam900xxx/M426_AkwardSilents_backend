package server

import (
	"AkwardSilents/pkg/service"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
)

func (s *Server) readLoopAccount(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF { // client disconnected
				break
			}
			fmt.Println("read error:", err) // other read error
			continue
		}
		msg := buf[:n]
		_, err = ws.Write([]byte(service.MessageHandlerAccount(msg))) //Process the received message
		if err != nil {
			return
		}
	}
}

func (s *Server) readLoopChat(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF { // client disconnected
				break
			}
			fmt.Println("read error:", err) // other read error
			continue
		}
		msg := buf[:n]
		fmt.Println("Received message:", string(msg))
	}
}
