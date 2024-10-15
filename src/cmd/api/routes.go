package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", http.NotFound) // Catch-all route
	mux.HandleFunc("/api/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/api/v1/version", app.versioninfo)

	return mux
}
