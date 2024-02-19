package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port string
}

type application struct {
	templateCache map[string]*template.Template
}

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

	// TEMPLATE CACHE
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		log.Fatal(err)
	}

	// CREATE SERVER
	app := &application{
		templateCache: templateCache,
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.port),
		Handler: app.routes(),
	}

	// START SERVER
	log.Printf("server started on %v\n", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
