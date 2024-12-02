package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	headers := map[string][]string(r.Header)
	if err := json.NewEncoder(w).Encode(headers); err != nil {
		http.Error(w, "Failed to encode headers", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", echoHandler)
	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
