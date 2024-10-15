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

	// mux.HandleFunc("/", Home)
	// mux.HandleFunc("/helloworld", HelloWorld)
	// mux.HandleFunc("/api/v1/healthcheck", app.HealthCheck)
	// mux.HandleFunc("/api/v1/version", app.VersionInfo)

	addr := fmt.Sprintf(":%d", cfg.port)

	logger.Printf("Starting %s server on %s", cfg.env, addr)
	err := http.ListenAndServe(addr, mux) // Locally Scoped Custom ServeMux.
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
