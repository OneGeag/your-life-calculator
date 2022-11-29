package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
	env  string
}

type appilcation struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	cfg.port = 6060
	cfg.env = "development"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &appilcation{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", app.ping)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: mux,
	}

	logger.Printf("Starting %s server on %s\n", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
