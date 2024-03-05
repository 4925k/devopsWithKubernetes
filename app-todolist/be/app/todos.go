package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	stmt := `INSERT INTO todolist (description) VALUES ($1)`

	_, err = db.Exec(stmt, input.Todo)
	if err != nil {
		log.Println(err)
		w.Write([]byte("database error"))
		return
	}

	out := fmt.Sprintf("todo created: %s", input.Todo)
	w.Write([]byte(out))
}

func getAllTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := GetAllTodo()
	if err != nil {
		w.Write([]byte("database error"))
		log.Println(err)
		return
	}

	out, err := json.Marshal(todos)
	if err != nil {
		w.Write([]byte("internal server error"))
		return
	}

	w.Write(out)
}
