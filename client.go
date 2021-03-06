package main

import (
	"github.com/gorilla/websocket"
	"time"
)

// client represent one chatting user

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
			// if avatarUrl, ok := c.userData["avatar_url"]; ok {
			// 	msg.AvatarURL = avatarUrl.(string)
			// }
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
	c.socket.Close()
}
