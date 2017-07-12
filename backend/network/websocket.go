package network

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Zenika/MARIE/backend/apiai"

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

// Handle the requests on websockets
func Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		var req = Request{}
		if err := json.Unmarshal(p, &req); err != nil {
			panic(err)
		}
		if req.Type == "speech" {
			conn.WriteJSON(apiai.Analyze(req.Message))
		}
	}
}
