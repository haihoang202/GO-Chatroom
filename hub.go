package main

type Hub struct {
	clients map[int] *Client
	broadcast chan[] byte
	join chan *Client
	exit chan *Client
	id int
}

func newHub() *Hub {
	hub := &Hub{
		clients:	make(map[int] *Client),
		broadcast:	make(chan [] byte),
		join:		make(chan *Client),
		exit: 		make(chan *Client),
		id:			0}

	return hub

}

func (hub *Hub) run() {
	for {
		select{
		case client := <- hub.join:
			client.clientID = hub.id
			hub.clients[hub.id] = client
			hub.id+= 1
			println("Client joined, ID: ", hub.id-1)
		case client := <- hub.exit:
			println(client)
		case msg := <- hub.broadcast:
			println("Hub receving and preparing to send the message",string(msg))
			for _,client := range hub.clients {
				client.send <- msg
			}

		}
	}
}