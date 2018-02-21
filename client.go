package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub   *Hub
	conn  *websocket.Conn
	send  chan Message
	id    string
	color string
}

type Message struct {
	Type  string `json:"type"`
	From  string `json:"senderID"`
	To    string `json:"receiverID"`
	Data  string `json:"data"`
	Color string `json:"color"`
}

var upgrader = websocket.Upgrader{}

func serveClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
	var conn, err = upgrader.Upgrade(w, r, nil)

	if err != nil {
		println("Error in upgrading http to websocket. Check log for more info")
		log.Println(err)
		return
	}

	var initialsMsg Message
	err = conn.ReadJSON(&initialsMsg)
	if err != nil {
		println("Error! Something wrong")
		conn.Close()
		return
	}

	println(initialsMsg.From, "joined the chatroom")

	client := &Client{
		hub:   hub,
		conn:  conn,
		send:  make(chan Message),
		id:    initialsMsg.From,
		color: initialsMsg.Color}

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
			client.hub.exit <- client
			client.conn.Close()
			break
		}

		println("Reading from user", string(msg.From), "type: ", string(msg.Type))
		switch msg.Type {
		case "broadcast":
			client.hub.broadcast <- msg
		case "private":
			client.hub.private <- msg
		}
	}
}

func (client *Client) writeMsg() {
	for {
		select {
		case msg := <-client.send:
			println("Receving message: ", msg.Data)
			client.conn.WriteJSON(Message{
				Type:  msg.Type,
				From:  msg.From,
				To:    msg.To,
				Data:  msg.Data,
				Color: msg.Color,
			})
		}
	}
}
