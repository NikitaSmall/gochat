package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":3000", "The addr of the application")
	flag.Parse()
	r := newRoom()
	mainTemplate := &templateHandler{filename: "chat.html"}

	http.HandleFunc("/", mainTemplate.ServeHTTP)
	http.HandleFunc("/room", r.serveHTTP)

	go r.run()

	log.Println("Server starting on ", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
