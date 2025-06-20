package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))

	const address = ":8080"
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Failed to start server on %s: %v", address, err)
	}
}
