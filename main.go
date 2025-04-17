package main

import (
	"log"
)

func main() {
	server := SetupServer()

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
