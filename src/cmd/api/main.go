package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const versionNumber = "1.0.1"

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
	flag.StringVar(&cfg.env, "env", "Development", "Environment the server is running in")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/healthcheck", healthcheck)
	mux.HandleFunc("/api/v1/version", version)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	logger.Printf("Starting %s server on %s", cfg.env, addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Status: Healthy")
}

func version(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Environment: %s\n", "Development")
	fmt.Fprintf(w, "Version: %s\n", versionNumber)
}
