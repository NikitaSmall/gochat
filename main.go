package main

import (
	"log"
	"net/http"
)

var mainTemplate = &templateHandler{filename: "chat.html"}

func main() {
	r := newRoom()

	http.HandleFunc("/", mainTemplate.ServeHTTP)
	http.HandleFunc("/room", r.serveHTTP)

	go r.run()

	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
