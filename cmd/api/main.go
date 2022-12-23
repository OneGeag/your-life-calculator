package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	cfg.port = 6060
	cfg.env = "development"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: mux,
	}

	logger.Printf("Starting %s server on %s\n", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
