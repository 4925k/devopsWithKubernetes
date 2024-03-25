package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("app started")

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

	r.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hash finder app"))
	}))
	r.HandleFunc("/hash/status", getHash)

	return r

}

// getHash will get the current hash and pingpong status from the current shared file
func getHash(w http.ResponseWriter, r *http.Request) {
	// GET HASH
	f, err := os.OpenFile("/usr/src/app/files/hash.txt", os.O_RDONLY, 0644)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer f.Close()

	bd, err := io.ReadAll(f)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	// GET PONG

	// from file
	// f2, err := os.OpenFile("/usr/src/app/files/pingpong.txt", os.O_RDONLY, 0644)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// b2, err := io.ReadAll(f2)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// from endpoint
	pp, err := http.Get("http://pingpong-svc:2347/pingpong")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer pp.Body.Close()

	ppbd, err := io.ReadAll(pp.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	bd = append(bd, ppbd...)

	greet := os.Getenv("GREET")
	if greet == "" {
		greet = "env empty"
	}

	bd = append(bd, []byte(greet)...)

	w.Write(bd)
}
