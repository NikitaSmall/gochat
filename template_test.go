package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()

	mainTemplate := &templateHandler{filename: "chat.html"}
	mainTemplate.ServeHTTP(w, req)

	if w.Code != 200 || !strings.Contains(w.Body.String(), "Chat!") {
		t.Errorf("Expected code '200' and body 'Chat!', got code '%v', body '%v'", w.Code, w.Body)
	}
}
