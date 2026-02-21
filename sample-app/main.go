package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, CI/CD World! Version 1.0.0")
	})
	fmt.Println("Server starting on port 8081...")
	http.ListenAndServe(":8081", nil)
}
