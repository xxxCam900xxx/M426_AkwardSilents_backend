package handler

import (
	"encoding/json"
	"fmt"
	"AkwardSilents/pkg/service/functions"
    "strings"
)

// Define a struct to hold the JSON data
type message struct {
	Typ     string            `json:"typ"` // sendmessage, getmessage, getoverview, getmembers  vt
	Content	map[string]string `json:"content"`
}

func Warumdasnicht(msg []byte) {
	fmt.Println("Message handler bekommen")

	// Create an empty Message object
	var message message
	// Unmarshal the JSON string into the Message object
	msg = []byte(strings.ReplaceAll(string(msg), "'", "\""))
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(message.Typ)
		fmt.Println(message.Content)
		switch message.Typ {
		case "login":
			functions.Login(message.Content);
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