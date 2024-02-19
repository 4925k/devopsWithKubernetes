package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var todolist []string

func createTodo(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Todo string `json:"todo"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Write([]byte("error while reading body"))
		return
	}

	todolist = append(todolist, input.Todo)

	out := fmt.Sprintf("todo created: %s", input.Todo)
	w.Write([]byte(out))
}

func getAllTodo(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(todolist)
	if err != nil {
		w.Write([]byte("internal server error"))
		return
	}

	w.Write(out)
}
