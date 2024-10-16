package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/helloworld", HelloWorld)

	mux.HandleFunc("/api/v1/healthcheck", app.HealthCheck)
	mux.HandleFunc("/api/v1/version", app.VersionInfo)

	mux.HandleFunc("/api/v1/books", app.getBooksHandler)

	return mux
}
