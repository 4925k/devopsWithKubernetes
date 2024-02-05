package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("/hello-world", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello user"))
	}))
	r.HandleFunc("/hash/get", app.hashHandler)

	return r
}
