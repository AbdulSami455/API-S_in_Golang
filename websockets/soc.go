package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		if messageType == websocket.TextMessage {
			message := string(p)
			fmt.Printf("Received: %s\n", message)

			// You can handle the received message here

			// Send a response back to the client
			responseMessage := []byte("Received: " + message)
			if err := conn.WriteMessage(websocket.TextMessage, responseMessage); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocketConnection)
	http.Handle("/", http.FileServer(http.Dir("."))) // Serve static files from the current directory

	port := ":8080"
	fmt.Printf("Server started on %s\n", port)
	http.ListenAndServe(port, nil)
}
