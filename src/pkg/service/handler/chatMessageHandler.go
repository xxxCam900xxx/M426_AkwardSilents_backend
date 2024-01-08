package handler

import (
	"encoding/json"
	"fmt"
)

// Define a struct to hold the JSON data
type message struct {
	Typ     string            `json:"typ"` // sendmessage, getmessage, getoverview, getmembers  vt
	Content map[string]string `json:"content"`
}

func MessageHandlerChats(msg []byte) {
	// Create an empty Message object
	var message message
	// Unmarshal the JSON string into the Message object
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		switch message.Typ {
		case "login":
			login(message.Content);
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