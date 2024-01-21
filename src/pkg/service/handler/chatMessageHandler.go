package handler

import (
	"encoding/json"
	"fmt"
	"AkwardSilents/pkg/service/functions"
    "strings"
	"golang.org/x/net/websocket"
)

// Define a struct to hold the JSON data
type message struct {
	Typ     string            `json:"typ"` // sendmessage, getmessage, getoverview, getmembers  vt
	Content	map[string]string `json:"content"`
}

func Chat(msg []byte, ws *websocket.Conn, name *string) {
	var message message
	// Unmarshal the JSON string into the Message object
	msg = []byte(strings.ReplaceAll(string(msg), "'", "\""))
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		switch message.Typ {
		case "login":
			_, err = ws.Write([]byte(functions.Login(message.Content, name)))
		case "sendmessage":
			//sendmessage(message.Content)
		case "getmessage":
			//getmessage(message.Content)
		case "getoverview":
			//getoverview(message.Content)
		case "getmembers":
			//getmembers(message.Content)
		}
	}
}