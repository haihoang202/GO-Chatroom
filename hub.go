package main

type Hub struct {
	clients map[int] *Client
	broadcast chan Message
	join chan *Client
	exit chan *Client
	id int
}

func newHub() *Hub {
	hub := &Hub{
		clients:	make(map[int] *Client),
		broadcast:	make(chan Message),
		join:		make(chan *Client),
		exit: 		make(chan *Client),
		id: 		0}

	return hub

}

func (hub *Hub) run() {
	for {
		select{
		case client := <- hub.join:
			hub.clients[hub.id] = client
			hub.id += 1
			// hub.clients = append(hub.clients,client)
			// println("Client joined, ID: ", client.clientID)
		case client := <- hub.exit:
			println(client)
		case msg := <- hub.broadcast:
			println("Hub receving and preparing to send the message",msg.Data)
			for _,client := range hub.clients {
				client.send <- msg
			}

		}
	}
}