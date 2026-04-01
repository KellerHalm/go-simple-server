package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "OK")
}

func main() {
    const version = "1.1.0" 
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello! Server Version: %s", version)
    })
    http.HandleFunc("/health", healthHandler) 
    
    fmt.Printf("Starting server v%s on port :8080\n", version)
    http.ListenAndServe(":8080", nil)
}