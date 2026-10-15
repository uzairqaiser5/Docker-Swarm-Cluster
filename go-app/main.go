package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go running inside a minimal Docker container!")
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
