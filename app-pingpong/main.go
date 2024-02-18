package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
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

	// pingpong keeps a counter to the number of times
	// this endpoint has been called
	r.Handle("/pingpong", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// increase counter
		counter++

		// prep output
		out := fmt.Sprintf("pong %d", counter)

		// write to file
		os.WriteFile("/usr/src/app/files/pingpong.txt", []byte(out), fs.FileMode(os.O_CREATE))

		// send back to client
		w.Write([]byte(out))
	}))

	return r
}
