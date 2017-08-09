package nWS

// Hub represents a websocket client pool
type Hub struct {
	clients      map[*Client]bool
	broadcast    chan []byte
	addClient    chan *Client
	removeClient chan *Client
}

// initialize new hub
var hub = Hub{
	broadcast:    make(chan []byte),
	addClient:    make(chan *Client),
	removeClient: make(chan *Client),
	clients:      make(map[*Client]bool),
}

// Runs forever as a goroutine
func (hub *Hub) start() {
	for {
		// one of these fires when a channel
		// receives data
		select {
		case conn := <-hub.addClient:
			// add a new client
			hub.clients[conn] = true
		case conn := <-hub.removeClient:
			// remove a client
			if _, ok := hub.clients[conn]; ok {
				delete(hub.clients, conn)
				close(conn.send)
			}
		case message := <-hub.broadcast:
			// broadcast a message to all clients
			for conn := range hub.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(hub.clients, conn)
				}
			}
		}
	}
}
