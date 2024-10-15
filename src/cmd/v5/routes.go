package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", http.NotFound) // Catch-all route
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/helloworld", HelloWorld)

	mux.HandleFunc("/api/v1/healthcheck", app.HealthCheck)
	mux.HandleFunc("/api/v1/version", app.VersionInfo)

	mux.HandleFunc("/api/v1/books", app.getBooksHandler)
	mux.HandleFunc("/api/v1/books", app.createBookHandler)

	return mux
}
