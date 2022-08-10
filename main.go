package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())

	if err := http.ListenAndServe(":5555", server); err != nil {
		log.Fatalf("Could not listen on port %v", err)
	}
}
