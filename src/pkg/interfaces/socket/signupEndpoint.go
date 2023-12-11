package socket

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func signupEndpoint() {
	server := NewServer()
	http.Handle("/account", websocket.Handler(server.handleWS))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting signup server:", err)
	}
}
