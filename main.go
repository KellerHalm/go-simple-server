package main

import (
	"fmt"
	"net/http"
)

const version = "1.0.0"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! Server Version: %s", version)
	})

	fmt.Printf("Starting server v%s on port :8080\n", version)
	http.ListenAndServe(":8080", nil)
}