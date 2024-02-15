package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port string
}

type application struct{}

func main() {
	// GET APPLICATION CONFIG
	cfg := config{
		port: os.Getenv("PORT"),
	}

	// FETCH DAILY IMAGE
	getDailyImage()

	// CHANGE IMAGE EVERY 24 hourse
	go func() {
		for range time.Tick(time.Hour * 24) {
			getDailyImage()
		}
	}()

	// CREATE SERVER
	var app application
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.port),
		Handler: app.routes(),
	}

	// START SERVER
	log.Printf("server started on %v\n", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
