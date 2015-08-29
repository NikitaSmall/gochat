package main

import (
	"log"
)

type message struct {
	Message []byte
}

type messages []message

func (m *messages) addMessage(message message) {
	log.Println("message %s was stored", message.Message)
	*m = append(*m, message)
}

func newMessagePull() *messages {
	m := make(messages, 0, 0)
	return &m
}
