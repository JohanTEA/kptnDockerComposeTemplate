package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// Define build variables here. Variables are set at build time with '-X main.<variable>'
var version string

func main() {
    // Print the build version
    fmt.Println("App Version:", version)
    
    // Define the endpoint
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
    })

    // Start server on port 8080
    fmt.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

