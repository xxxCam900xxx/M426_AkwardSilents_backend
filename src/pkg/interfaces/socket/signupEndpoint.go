package socket

import (
	"AkwardSilents/pkg/service/server"
	"golang.org/x/net/websocket"
	"net/http"
)

func SignupEndpoint() error {
	s := server.NewServer()
	http.Handle("/account", websocket.Handler(s.HandleWSAccount))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return err
	}
	return nil
}
