package main

import (
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":3000", "The addr of the application")
	flag.Parse()

	// set up gomniauth
	gomniauth.SetSecurityKey("some long key")
	gomniauth.WithProviders(
		facebook.New("key", "secret",
			"http://localhost:3000/auth/callback/facebook"),
		github.New("key", "secret",
			"http://localhost:3000/auth/callback/github"),
		google.New("89198476902-sp9n6ukjk37ccc2vuh11chkga865sidq.apps.googleusercontent.com", "GeQA1wriK2DdN_tTg8Sv1x9r",
			"http://localhost:3000/auth/callback/google"),
	)

	r := newRoom()
	http.HandleFunc("/room", r.serveHTTP)

	mainTemplate := &templateHandler{filename: "chat.html", messages: r.Messages}
	loginTemplate := &templateHandler{filename: "login.html"}

	http.HandleFunc("/auth/", LoginHandler)
	http.HandleFunc("/chat", MustAuth(mainTemplate).ServeHTTP)
	http.HandleFunc("/login", loginTemplate.ServeHTTP)

	go r.run()

	log.Println("Server starting on ", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
