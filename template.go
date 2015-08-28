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
}

// ServeHTTP handles http request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}
