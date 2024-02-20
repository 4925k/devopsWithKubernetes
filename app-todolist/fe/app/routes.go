package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", app.home)
	r.Handle("/hello-world", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello user"))
	}))

	r.HandleFunc("/hash/get", app.hashHandler)
	r.HandleFunc("POST /todos/create", app.createTodo)

	fs := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/", http.StripPrefix("/static", fs))

	return r
}
