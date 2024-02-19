package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3001),
		Handler: routes(),
	}

	log.Print("starting server")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("server started at %s", srv.Addr)
}
