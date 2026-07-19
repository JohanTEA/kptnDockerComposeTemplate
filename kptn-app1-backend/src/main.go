package main

import (
    "encoding/json"
    "fmt"
    "strconv"
    "log"
    "net/http"
)

// Define build variables here. Variables are set at build time with '-X main.<variable>'
var version string

func main() {
    visitors := 0
    
    // Print the build version
    fmt.Println("App Version:", version)
    
    // Define endpoints
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        // Allow all origins (use specific domain in production)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
	// Handle request
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	visitors++
    	visitorsStr := strconv.Itoa(visitors)
    	
    	// Allow all origins (use specific domain in production)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
	// Handle request
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"visitors": visitorsStr})
    })
 
    // Start server
    fmt.Println("Server running on :80")
    log.Fatal( http.ListenAndServe(":80", nil) )
}

