package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Message Define a struct to hold the JSON data
type register struct {
	Displayname string `json:"displayname"`
	Username    string `json:"username"`
}

func MessageHandlerAccount(msg []byte) string {
	// Create an empty Message object
	var message register
	// Unmarshal the JSON string into the Message object
	msg = []byte(strings.ReplaceAll(string(msg), "'", "\""))
	fmt.Println(string(msg))
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("Errorrrrr:", err)
	} else {
		// Print the Message objects
		fmt.Printf("Message: %+v\n", message.Username)
		fmt.Printf("Message: %+v\n", message.Displayname)

		if true { //TODO: username validation (username exists already in database?).
			return createToken() //TODO: Return Token for signup.
		}
	}

	return ""
}

func createToken() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var hashlength = 20
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
	hash := make([]rune, hashlength)
	for i := range hash {
		hash[i] = letters[rand.Intn(len(letters))]
	}
	return string(hash)
}
