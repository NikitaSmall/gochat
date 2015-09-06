package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ErrorNoAvatarURL is an error that fires in case client hasn't avatarURL
var ErrorNoAvatarURL = errors.New("Chat: unable to get avatarURL")

// Avatar interface represents types capable of representing of user profile pictures
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar
type AuthAvatar struct{}

type GravatarAvatar struct{}

// UseAuthAvatar
var UseAuthAvatar AuthAvatar

var UseGravatarAvatar GravatarAvatar

// GetAvatarURL
func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrorNoAvatarURL
}

func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x",
				m.Sum(nil)), nil
		}
	}
	return "", ErrorNoAvatarURL
}
