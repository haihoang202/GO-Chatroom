package main

type Hub struct {
	clients   map[string]*Client
	broadcast chan Message
	private   chan Message
	join      chan *Client
	exit      chan *Client
}

func newHub() *Hub {
	hub := &Hub{
		clients:   make(map[string]*Client),
		broadcast: make(chan Message),
		private:   make(chan Message),
		join:      make(chan *Client),
		exit:      make(chan *Client)}

	return hub

}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.join:
			if _, ok := hub.clients[client.id]; ok {
				println("Duplicated id!")
			} else {
				hub.clients[client.id] = client
			}
		case client := <-hub.exit:
			println(client)
			delete(hub.clients, client.id)
		case msg := <-hub.broadcast:
			println("[BROADCASTING] message", msg.Data)
			for _, client := range hub.clients {
				client.send <- msg
			}
		case msg := <-hub.private:
			println("[PRIVATE] from", msg.From, "to", msg.To)
			hub.clients[msg.To].send <- msg
			hub.clients[msg.From].send <- msg

		}
	}
}
