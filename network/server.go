package network

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"textEditor/util"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity, but consider security in production
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Server: Error upgrading connection: %v", err)
		return
	}
	defer ws.Close()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Server: Error reading message: %v", err)
			break
		}
		log.Printf("Server: Received: %s", message)
		if messageType == websocket.TextMessage {
			err = ws.WriteMessage(websocket.PongMessage, []byte(util.UserName))
			if err != nil {
				log.Printf("Server: Error writing pong message: %v", err)
				break
			}
			continue
		}
		//TODO: implement adding text massages to the text area
		err = ws.WriteMessage(messageType, []byte("Echo: "+string(message)))
		if err != nil {
			log.Printf("Server: Error writing message: %v", err)
			break
		}
	}
}

func StartWebSocketServer() {
	http.HandleFunc("/ws", handleConnections)
	log.Printf("WebSocket server starting on :%v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
