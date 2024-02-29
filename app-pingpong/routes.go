package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func routes() *http.ServeMux {
	r := http.NewServeMux()

	// pingpong keeps a counter to the number of times
	// this endpoint has been called
	r.Handle("/pingpong", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get counter
		currentCount := GetCount()
		log.Println("count fetched", currentCount)

		// prep output
		out := fmt.Sprintf("pong %d", currentCount)

		// write to file
		os.WriteFile("/usr/src/app/files/pingpong.txt", []byte(out), fs.FileMode(os.O_CREATE))

		// increase counter
		IncreaseCounter()
		log.Println("count increased")

		// send back to client
		w.Write([]byte(out))
	}))

	return r
}
