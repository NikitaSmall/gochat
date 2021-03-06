package main

import (
	"log"
	"time"
)

type message struct {
	Message   string
	When      time.Time
	Name      string
	AvatarURL string
}

type messages []message

func (m *messages) addMessage(message message) {
	log.Printf("message %s was stored", string(message.Message))
	*m = append(*m, message)
}

func (m *messages) toString() []string {
	var strings []string
	for _, message := range *m {
		strings = append(strings, string(message.Message))
	}

	return strings
}

func newMessagePull() *messages {
	m := make(messages, 0, 0)
	return &m
}
