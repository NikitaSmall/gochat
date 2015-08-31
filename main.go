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
	http.HandleFunc("/room", r.serveHTTP)

	mainTemplate := &templateHandler{filename: "chat.html", messages: r.Messages}
	http.HandleFunc("/chat", MustAuth(mainTemplate).ServeHTTP)

	go r.run()

	log.Println("Server starting on ", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
