package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("*html")
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	tmpl.Execute(w, nil)

}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade the connection to a websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	if err := ws.WriteMessage(1, []byte("hi client!")); err != nil {
		log.Println(err)
		return
	}
	// helpful log statement to show connection
	log.Println("client connected...")
	reader(ws)

}

// listen indefinitely for new messages coming
// through on our WebSocket connection.
//
// define a reader which will listen for
// new messages being sent to our WebSocket endpoint
func reader(conn *websocket.Conn) {
	msg := []byte("Wow! message reached to golang server, fantanstic. the message is : ")

	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out that msg
		fmt.Println("message type is :", messageType)
		fmt.Println("message is :", string(p))

		msg = append(msg, p...)

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Lets begin")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
