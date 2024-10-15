// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"
// )

// const versionNumber = "1.0.1"

// type serverConfig struct {
// 	port int
// 	env  string
// }

// type application struct {
// 	config serverConfig
// 	logger *log.Logger
// }

// func main() {
// 	var cfg serverConfig

// 	flag.IntVar(&cfg.port, "port", 8080, "Port to run the server on")
// 	flag.StringVar(&cfg.env, "env", "Development", "Environment the server is running in")
// 	flag.Parse()

// 	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

// 	app := &application{
// 		config: cfg,
// 		logger: logger,
// 	}

// 	addr := fmt.Sprintf(":%d", cfg.port)

// 	srv := &http.Server{
// 		Addr:         addr,
// 		Handler:      app.routes(),
// 		IdleTimeout:  time.Minute,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 	}

// 	logger.Printf("Starting %s server on %s", cfg.env, addr)
// 	err := srv.ListenAndServe()
// 	if err != nil {
// 		logger.Printf("Error starting server: %v\n", err)
// 		os.Exit(1)
// 	}
// }

// // func healthcheck(w http.ResponseWriter, r *http.Request) {
// // 	if r.Method != http.MethodGet {
// // 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// // 		return
// // 	}

// // 	fmt.Fprintln(w, "Status: Healthy")
// // }

// // func versioninfo(w http.ResponseWriter, r *http.Request) {
// // 	if r.Method != http.MethodGet {
// // 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// // 		return
// // 	}

// // 	fmt.Fprintf(w, "Environment: %s\n", "Development")
// // 	fmt.Fprintf(w, "Version: %s\n", versionNumber)
// // }
