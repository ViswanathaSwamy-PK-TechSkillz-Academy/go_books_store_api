package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/healthcheck", healthcheck)
	http.HandleFunc("/v1/version", version)

	fmt.Println("Starting ... Server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: Healthy")
}

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Environment: %s\n", "Development")
	fmt.Fprintf(w, "Version: %s\n", "1.0.1")
}
