package main

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	Once     sync.Once
	filename string
	templ    *template.Template
	messages *messages
}

// ServeHTTP handles http request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	// t.templ.Funcs(FuncMap{"messages": func() *messages { return t.messages }})
	t.templ.Execute(w, map[string]interface{}{
		"r":        r,
		"messages": t.messages.toString(),
	})
}
