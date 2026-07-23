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
	requestCounter := 0

	// Print the build version
	fmt.Println("App Version:", version)

	// Define endpoints
	//   All API.s used by the frontend or externally must start with '/api'.
	//   The revese-proxy is set-up to route to this app for '/api' on the same root adress as frontend.
	//   This is to comply with CORS.
	//   Internal API.s do not need to be behind '/api'.
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		requestCounter++

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":   "ok",
			"requests": strconv.Itoa(requestCounter),
		})
	})

	// Start server
	fmt.Println("Server running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
