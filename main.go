package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
      <html>
        <head>
          <title>Chat!</title>
        </head>
        <body>
          Let's chat!
        </body>
      </html>
    `))
}

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
