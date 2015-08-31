package main

import (
	"net/http"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		// not authorized
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		// other error
		panic(err.Error())
	} else {
		// success - call next handler
		h.next.ServeHTTP(w, r)
	}
}

// MustAuth is a function to check and serve auth before chatting
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
