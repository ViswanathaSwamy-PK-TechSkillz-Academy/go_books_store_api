package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Welcome to Go Lang World, you've requested: %s\n", r.URL.Path)
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
