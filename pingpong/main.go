package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int

func main() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "4445"),
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("pingpong started on %s", srv.Addr)
}

func routes() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("/pingpong", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter++
		out := fmt.Sprintf("pong %d", counter)
		w.Write([]byte(out))
	}))

	return r
}