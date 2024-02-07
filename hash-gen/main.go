package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var currentStatus string

func main() {
	// generate a hash and time every 5 second
	go func() {
		for {
			h, err := bcrypt.GenerateFromPassword([]byte(time.Now().String()), 8)
			if err != nil {
				log.Fatal(err)
				return
			}

			currentStatus = fmt.Sprintf("%s %s\n", time.Now().String(), h)

			fmt.Println(currentStatus)

			time.Sleep(time.Second * 5)
		}
	}()

	srv := &http.Server{
		Addr:    ":4444",
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func routes() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("/hash/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(currentStatus))
	}))

	return r
}
