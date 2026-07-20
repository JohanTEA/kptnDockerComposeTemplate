package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define build variables here. Variables are set at build time with '-X main.<variable>'
var version string

func main() {
	visitors := 0

	// Print the build version
	fmt.Println("App Version:", version)

	// Define endpoints
	//   All endpoints must start with '/api' for the routing setup to work and to be able to
	//   publish the API on the same root-address as the frontend to comply with CORS.
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		visitors++
		requestsStr := strconv.Itoa(visitors)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"requests": requestsStr})
	})

	// Start server
	fmt.Println("Server running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
