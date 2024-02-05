package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port string
}

type application struct{}

func main() {
	cfg := config{
		port: os.Getenv("PORT"),
	}

	var app application
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.port),
		Handler: app.routes(),
	}

	log.Printf("server started on %v\n", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
