package tools

var users = make(map[string])

func addName(name string, ws *websocket.Conn) {
    if name != "" {
        users[name] = &User{Name: name, Conn: ws}
    }
}

func sendMessage(name string, message string) string {
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

func removeName(name string){
    delete(users, name)
}