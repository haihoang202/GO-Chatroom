package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"

)

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan[] byte
	clientID int
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
		send: make(chan[] byte, 256),
		clientID: 0}

	client.hub.join <- client

	go client.readMsg()
	go client.writeMsg()
}

func (client *Client) readMsg() {
	for {
		// var newmsg Message
		// _, msg, err := client.conn.ReadJSON(&newmsg)
		_, msg, err := client.conn.ReadMessage()
		if err != nil {
			println("Client left")
			client.conn.Close()
			break
		}

		println("Reading from user and broadcasting: ", string(msg))
		client.hub.broadcast <- msg
	}
}

func (client *Client) writeMsg() {
	for {
		select{
		case msg := <-client.send:
			println("Receving message: ", string(msg))
			client.conn.WriteJSON(Message{
				Client:	client.clientID,
				Data: 	string(msg),
			})
		}	
	}
}

type Message struct {
	Client 	int		`json:"clientID"`
	Data	string	`json:"data"`
}