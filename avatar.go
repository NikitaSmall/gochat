package main

import (
	"errors"
)

// ErrorNoAvatarURL is an error that fires in case client hasn't avatarURL
var ErrorNoAvatarURL = errors.New("Chat: unable to get avatarURL")

// Avatar interface represents types capable of representing of user profile pictures
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client
	GetAvatarURL(c *client) (string error)
}

// AuthAvatar
type AuthAvatar struct{}

// UseAuthAvatar
var UseAuthAvatar AuthAvatar

// GetAvatarURL
func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrorNoAvatarURL
}
