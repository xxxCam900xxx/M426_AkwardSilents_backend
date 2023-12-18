package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct to hold the JSON data
type Message struct {
    Typ string `json:"typ"` // sendmessage, getmessage, getoverview, getmembers
    Content map[string]string `json:"content"`
}

func MessageHandlerChats(msg []byte, name string) {
	if name != "" {
		// Create an empty Message object
		var message Message
		// Unmarshal the JSON string into the Message object
		err := json.Unmarshal(msg, &message)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			// Print the Message object
			switch message.Typ {
			case "sendmessage":
				sendmessage(message.Content)
			case "getmessage":
				getmessage(message.Content)
			case "getoverview":
				getoverview(message.Content)
			case "getmembers":
				getmembers(message.Content)
			default:
				fmt.Println(message)

				}			
		}
	}
}

func main() {
	// A sample JSON string
	jsonString := []byte(`{"typ":"getmessage","content":{"text":"wqe"}}`)
	// Call the function with the JSON string and a name
	MessageHandlerChats(jsonString, "Bob")
}

