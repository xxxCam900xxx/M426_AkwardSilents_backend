package socket

import (
	"AkwardSilents/pkg/service/server"
	"golang.org/x/net/websocket"
	"net/http"
)

func ChatEndpoint() error {
	s := server.NewServer()
	http.Handle("/chat", websocket.Handler(s.HandleWSChat))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}
	return nil
}
