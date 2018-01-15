# GO-Chatroom
This is a simple chatroom, implemented in GO using Websocket Gorilla package, VueJS and Bulma in the front end.

To run:
<br>
go run *.go<br>
open index.html (couple of this tabs for interacting)
<br>
<br>
<br>
Back end:<br>
Have 3 pieces: main, hub and client <br>
main.go: <br>
- initiate a hub where client can connect to<br>
- serve index file<br>
- create websocket that serve message in and out <br>
hub.go: <br>
- hub listens to 3 activities:<br>
+ If client joins: add to a channel to manage <br>
+If client exits: remove client from a managemnet chanel <br>
+If a message is receive: broadcast to client in management chanel <br>
client.go: <br>
- update current http connection to websocket protocol <br>
- read message sent from current serving client and send to the hub for broadcasting<br>
- receive message broadcasted from the hub and display to the chat box <br>

Continue...
