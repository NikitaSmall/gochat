package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type room struct {
	forward  chan []byte
	join     chan *client
	leave    chan *client
	clients  map[*client]bool
	Messages *messages
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func newRoom() *room {
	return &room{
		forward:  make(chan []byte),
		join:     make(chan *client),
		leave:    make(chan *client),
		clients:  make(map[*client]bool),
		Messages: newMessagePull(),
	}
}

func (r *room) serveHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
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
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward messages to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send the messages
					message := message{Message: msg}
					r.Messages.addMessage(message)
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
