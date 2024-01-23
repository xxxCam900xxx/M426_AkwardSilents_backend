package service

import (
	"AkwardSilents/pkg/service/functions"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"golang.org/x/net/websocket"
)

// Message Define a struct to hold the JSON data
type register struct {
	Displayname string `json:"displayname"`
	Username    string `json:"username"`
}

func init() {
	/*err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Loaded .env file")
	*/
	log.Println("Log")
}

// Define a struct to hold the JSON data
type message struct {
	Typ     string            `json:"typ"` // sendmessage, getmessage, getoverview, getmembers  vt
	Content map[string]string `json:"content"`
}

func MessageHandlerChats(msg []byte, ws *websocket.Conn, name *string) {
	var message message
	// Unmarshal the JSON string into the Message object
	msg = []byte(strings.ReplaceAll(string(msg), "'", "\""))
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Typ von anfrage", message.Typ)
		switch message.Typ {
		case "login":
			_, err = ws.Write([]byte(functions.Login(message.Content, name, ws)))
		case "sendmessage":
			functions.Sendmessage(message.Content, *name)
		case "getmessage":
			_, err = ws.Write([]byte(functions.GetMessage(*name)))
		case "getoverview":
			_, err = ws.Write([]byte(functions.Overview(*name)))
		case "getmembers":
			//getmembers(message.Content)
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