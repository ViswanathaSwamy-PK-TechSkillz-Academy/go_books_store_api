package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.1"

type serverConfig struct {
	port int
	env  string
}

type application struct {
	config serverConfig
	logger *log.Logger
}

func main() {
	var cfg serverConfig

	flag.IntVar(&cfg.port, "port", 8080, "Port to run the server on")
	flag.StringVar(&cfg.env, "env", "Development", "Environment (Dev | Test | Stage | Prod) the server is running in")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	mux.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Welcome to Go Lang World, you've requested: %s\n", r.URL.Path)
	})

	mux.HandleFunc("/api/v1/healthcheck", healthcheck)
	mux.HandleFunc("/api/v1/version", app.versioninfo)

	addr := fmt.Sprintf(":%d", cfg.port)

	logger.Printf("Starting %s server on %s", cfg.env, addr)
	err := http.ListenAndServe(addr, mux) // Locally Scoped Custom ServeMux.
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

func (app *application) versioninfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Environment: %s\n", app.config.env)
	fmt.Fprintf(w, "Version: %s\n", version)
}
