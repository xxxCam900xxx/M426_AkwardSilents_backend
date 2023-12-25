1. Run Main.go (go run main.go)
2. Open Browser dev Tools
3. Open Konsole
4. Type:
    ````
    let socket = new WebSocket("ws://localhost:3000")
    ````
5. Enter
6. Type:
    ````
    socket.onmessage = (event) => { console.log("received form the server: ", event.data) }
    ````
7. Enter
8. Type:
    ````
    socket.send("hello from client")
    ````
9. Check the server console for Connection
10. Check the dev Console for "received form the server:  thx for your msg"
11. done.