package service

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Define a struct to hold the JSON data
type message struct {
	Typ     string            `json:"typ"` // sendmessage, getmessage, getoverview, getmembers  vt
	Content map[string]string `json:"content"`
}

// Message Define a struct to hold the JSON data
type register struct {
	Displayname string `json:"displayname"`
	Username    string `json:"username"`
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded .env file")
}

func MessageHandlerChats(msg []byte, name string) {
	if name != "" {
		// Create an empty Message object
		var message message
		// Unmarshal the JSON string into the Message object
		err := json.Unmarshal(msg, &message)
		if err != nil {
			fmt.Println("Unmarshal - Error:", err)
		} else {
			// Print the Message object
			switch message.Typ {
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
}

func MessageHandlerAccount(msg []byte) string {
	// Create an empty Message object
	var message register
	// Unmarshal the JSON string into the Message object
	msg = []byte(strings.ReplaceAll(string(msg), "'", "\""))
	fmt.Println(string(msg))
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("Unmarshal - Error:", err)
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
	letters := []rune(os.Getenv("KEYCONTENT"))
	hashLengthStr := os.Getenv("KEYLENGTH")

	hashLength, err := strconv.Atoi(hashLengthStr)
	if err != nil {
		hashLength = 10
	}

	hash := make([]rune, hashLength)

	for i := range hash {
		hash[i] = letters[rand.Intn(len(letters))]
	}

	return string(hash)
}
