package server

import (
	"AkwardSilents/pkg/service/handler"
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
		//TODO: Process the received message if needed
		//fmt.Println("Received message:", string(msg))

		_, err = ws.Write([]byte(handler.MessageHandlerAccount(msg)))
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
		//TODO: Process the received message if needed
		fmt.Println("Received message:", string(msg))
	}
}
