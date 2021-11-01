package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func getPrice( w http.ResponseWriter, r *http.Request ) {
	ws, _ := upgrader.Upgrade( w, r, nil )
	defer ws.Close()

	for {
		
		mt, message, _ := ws.ReadMessage()
		log.Printf("Message received: %s", message)

		// Response message
		_ = ws.WriteMessage(mt, message)
		log.Printf("Message sent: %s ", message)

	}
}

func main() {
	http.HandleFunc("/", getPrice)
	log.Fatal(http.ListenAndServe(":5555", nil))
}