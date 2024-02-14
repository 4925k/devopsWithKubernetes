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

	r.HandleFunc("/hash/status", getHash)

	return r

}

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
	f2, err := os.OpenFile("/usr/src/app/files/pingpong.txt", os.O_RDONLY, 0644)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	b2, err := io.ReadAll(f2)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	bd = append(bd, b2...)

	w.Write(bd)
}
