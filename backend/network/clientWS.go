package network

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// Client represents a websocket connection
type Client struct {
	ws *websocket.Conn
	// Hub passes broadcast messages to this channel
	send chan []byte
}

// Hub broadcasts a new message and this fires
func (c *Client) write() {
	// make sure to close the connection incase the loop exits
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.ws.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// New message received so pass it to the Hub
func (c *Client) read() {
	defer func() {
		hub.removeClient <- c
		c.ws.Close()
	}()

	for {
		_, p, err := c.ws.ReadMessage()
		if err != nil {
			hub.removeClient <- c
			c.ws.Close()
			break
		}
		var req = Request{}
		if err := json.Unmarshal(p, &req); err != nil {
			log.Println(err)
		} else {
			if req.Type == "speech" {
				c.ws.WriteJSON(Analyze(req.Message))
			}
		}
	}
}
