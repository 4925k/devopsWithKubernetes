package main

import "net/http"

func routes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /todos", createTodo)
	r.HandleFunc("GET /todos", getAllTodo)

	return r
}
