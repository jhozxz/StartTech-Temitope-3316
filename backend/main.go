package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Simple Health Check for AWS ALB
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Main API Endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mongoURI := os.Getenv("MONGO_URI")
		redisHost := os.Getenv("REDIS_HOST")
		
		fmt.Fprintf(w, "Welcome to StartTech API!\n")
		fmt.Fprintf(w, "Backend is running on: %s\n", os.Getenv("HOSTNAME"))
		
		// Simple debug info (Do not do this in real prod for security reasons)
		if mongoURI != "" {
			fmt.Fprintf(w, "MongoDB Configured: Yes\n")
		} else {
			fmt.Fprintf(w, "MongoDB Configured: No\n")
		}
		
		if redisHost != "" {
			fmt.Fprintf(w, "Redis Configured: Yes\n")
		}
	})

	port := ":8080"
	fmt.Println("Server starting on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
