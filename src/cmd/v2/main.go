package main

import (
	"fmt"
	"net/http"
)

// go run .

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Welcome to Go Lang World, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/api/v1/healthcheck", healthcheck)
	http.HandleFunc("/api/v1/version", versioninfo)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil) // DefaultServeMux.
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Status: Healthy")
}

func versioninfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Environment: %s\n", "Development")
	fmt.Fprintf(w, "Version: %s\n", "1.0.1")
}
