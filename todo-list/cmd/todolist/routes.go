package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	r := http.NewServeMux()

	r.Handle("/hello-world", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello user"))
	}))

	r.HandleFunc("/", app.home)
	r.HandleFunc("/hash/get", app.hashHandler)

	fs := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/", http.StripPrefix("/static", fs))

	return r
}
