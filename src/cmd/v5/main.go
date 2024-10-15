package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.env, addr)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
