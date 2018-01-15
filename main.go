package main 

import (
	"net/http"
	"log"
)

func main () {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"index.html")	
	})

	http.HandleFunc("/mess/ws", func(w http.ResponseWriter, r *http.Request){
		serveClient(hub,w,r)	
	})

	if err := http.ListenAndServe(":2002",nil); err != nil {
		log.Fatal("ListenAndServe Error: ",err)
	} 

}