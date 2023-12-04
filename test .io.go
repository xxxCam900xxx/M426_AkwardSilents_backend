package main

import (
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Ein Map, um die Namen der Benutzer zu speichern
	userNames := make(map[string]string)

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		// Wenn der Benutzer eine Nachricht mit dem Präfix /name sendet, wird diese als sein Name gespeichert
		so.On("chat message", func(msg string) {
			if msg[:5] == "/name" {
				userNames[so.Id()] = msg[6:]
				so.Emit("chat message", "Ihr Name wurde zu "+msg[6:]+" geändert.")
			} else {
				so.Emit("chat message", msg)
			}
		})

		// Wenn der Benutzer /name abruft, wird sein gespeicherter Name zurückgegeben
		so.On("/name", func() {
			name, ok := userNames[so.Id()]
			if ok {
				so.Emit("chat message", "Ihr Name ist "+name+".")
			} else {
				so.Emit("chat message", "Sie haben noch keinen Namen eingegeben.")
			}
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
			delete(userNames, so.Id())
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
