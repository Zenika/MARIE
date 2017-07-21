package network

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Request represent the what the request should be
type Request struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// BroadcastWS a message to all sockets
func BroadcastWS(m []byte) {
	hub.broadcast <- m
}

// StartHub starts the hub instance
func StartHub() {
	go hub.start()
}

// Handle the requests on websockets
func Handle(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := &Client{
		ws:   conn,
		send: make(chan []byte),
	}

	hub.addClient <- client

	go client.write()
	go client.read()
}
