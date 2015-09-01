package main

import (
	"github.com/gorilla/websocket"
	"github.com/nikitasmall/trace"
	"github.com/stretchr/objx"
	"log"
	"net/http"
	"os"
)

type room struct {
	forward  chan *message
	join     chan *client
	leave    chan *client
	clients  map[*client]bool
	Messages *messages
	tracer   trace.Tracer
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func newRoom() *room {
	return &room{
		forward:  make(chan *message),
		join:     make(chan *client),
		leave:    make(chan *client),
		clients:  make(map[*client]bool),
		Messages: newMessagePull(),
		tracer:   trace.New(os.Stdout),
	}
}

func (r *room) serveHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("Failed to get auth cookie:", err)
		return
	}

	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.read()
	client.write()
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
			r.tracer.Trace("New user joined")
		case client := <-r.leave:
			// leaving
			r.tracer.Trace("Client left")
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward messages to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the messages
					r.Messages.addMessage(*msg)
					r.tracer.Trace("Sending message")
				default:
					r.tracer.Trace("Clean up")
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
