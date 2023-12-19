package handler

import (
	"encoding/json"
	"fmt"
)

// Message Define a struct to hold the JSON data
type register struct {
	Displayname string `json:"displayname"`
	Username    string `json:"username"`
}

func MessageHandlerAccount(msg []byte) {
	// Create an empty Message object
	var message register
	// Unmarshal the JSON string into the Message object
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// Print the Message objects
		fmt.Printf("Message: %+v\n", message.Username)
		fmt.Printf("Message: %+v\n", message.Displayname)

		//TODO: username validation (username exists already in database?).
		//TODO: write to database.
		//TODO: Return Token for signup.
	}
}
