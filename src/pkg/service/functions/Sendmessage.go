package functions

import (
	db "AkwardSilents/pkg/service/dbfunctions"
	"AkwardSilents/pkg/tools"
	"encoding/json"
	"fmt"
	"time"
)

func Sendmessage(content map[string]string, userName string) string {
	if userName != "" {
		if CheckValues(content, []string{"Name", "Message"}) {
			sendName := content["Name"]
			message := content["Message"]
			messageStatus := 0
			timestampBase64 := time.Now().Unix()
			timestamp := int(timestampBase64)

			data := map[string]interface{}{
				"typ":      "message",
				"username": userName,
				"time":     timestamp,
				"message":  message,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
			}

			jsonString := string(jsonData)
			fmt.Println(jsonString)

			fmt.Println(sendName, message, messageStatus)
			if tools.SendMessageToUser(sendName, jsonString) {
				messageStatus = 1
			}

			if !getChatID(userName, sendName, message, messageStatus, timestamp, userName) {
				insertChat(userName, sendName)
				getChatID(userName, sendName, message, messageStatus, timestamp, userName)
			}
			fmt.Println("success send message")
			return "logout"
		}
		return "logout"
	}
	return "logout"
}

func getChatID(userName string, sendName string, message string, status int, time int, sender string) bool {
	defer db.DB.Close()
	db.InitDB()
    rows, err := db.DB.Query("SELECT id FROM Personchats WHERE user1 = ? AND user2 = ?  UNION SELECT id FROM Personchats WHERE user1 = ? AND user2 = ? ", userName, sendName, sendName, userName)
	if err != nil {
		panic(err)
	}
	
	// Ergebnisse ausgeben
	for rows.Next() {
		var id int
		err = rows.Scan(&id)    
		if err != nil {
			panic(err)
		}
        rows.Close()
		insertMessage(message, status, time, id, userName)
        return true
	}
	return false

}

func insertMessage(message string, status int, time int, idChat int, sender string) {
	defer db.DB.Close()
    db.InitDB()
	_, err := db.DB.Exec("INSERT INTO Message (message, status, time, id_Personchats, sender) values (?,?,?,?,?)", message, status, time, idChat, sender)
	if err != nil {
		panic(err)
	}
}

func insertChat(userName string, sendName string) {
	defer db.DB.Close()
    db.InitDB()
    _, err := db.DB.Exec("INSERT INTO Personchats (user1, user2) values (?,?)", userName, sendName)
	if err != nil {
		panic(err)
	}
}

func CheckValues(content map[string]string, values []string) bool {
	for _, value := range values {
		if _, ok := content[value]; !ok {
			return false
		}
	}
	return true
}
