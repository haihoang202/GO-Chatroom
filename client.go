package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"

)

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan Message
}

type Message struct {
	Client 	string	`json:"clientID"`
	Data	string	`json:"data"`
}

var upgrader = websocket.Upgrader{}

func serveClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
	var conn,err = upgrader.Upgrade(w,r,nil)

	if err != nil {
		println("Error in upgrading http to websocket. Check log for more info")
		log.Println(err)
		return
	}

	client := &Client{
		hub: hub,
		conn: conn,
		send: make(chan Message),}

	client.hub.join <- client

	go client.readMsg()
	go client.writeMsg()
}

func (client *Client) readMsg() {
	for {
		var msg Message
		err := client.conn.ReadJSON(&msg)
		if err != nil {
			println("Client left")
			client.conn.Close()
			break
		}

		println("Reading from user and broadcasting: ", string(msg.Data))
		client.hub.broadcast <- msg
	}
}

func (client *Client) writeMsg() {
	for {
		select{
		case msg := <-client.send:
			println("Receving message: ", msg.Data)
			client.conn.WriteJSON(Message{
				Client:	msg.Client,
				Data: 	msg.Data,
			})
		}	
	}
}
