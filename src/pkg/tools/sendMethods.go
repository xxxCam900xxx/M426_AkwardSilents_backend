package tools

import (
    "golang.org/x/net/websocket"
    "log"
)

type User struct {
    Name string
    Conn *websocket.Conn
}


var users = make(map[string] *User)

func AddName(name string, ws *websocket.Conn) {
    if name != "" {
        users[name] = &User{Name: name, Conn: ws}
    }
}

func SendMessageToUser(name string, message string) bool {
    user, ok := users[name]
    if ok {
        err := websocket.JSON.Send(user.Conn, message)
        if err != nil {
            log.Println("Error sending message:", err)
        }
        return true
    } else {
        return false
    } 
}

func RemoveName(name string){
    delete(users, name)
}