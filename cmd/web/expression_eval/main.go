package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {

	server := &http.Server{
		Addr:    "127.0.0.1" + port,
		Handler: routes(),
	}

	log.Printf("staring web server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
