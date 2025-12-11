// Production test - Dec 11, 2025
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

type Health struct {
	Healthy bool `json:"healthy"`
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Welcome to Go Test"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	health := Health{Healthy: true}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}
// Health check fix test - Thu Dec 11 14:51:42 PST 2025
