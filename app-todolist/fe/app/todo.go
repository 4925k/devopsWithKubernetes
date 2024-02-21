package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (app *application) createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	bd, err := json.Marshal(map[string]string{
		"todo": r.PostForm.Get("todo"),
	})
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return

	}

	res, err := http.Post("http://todolistbe:2344/todos", "application/json", bytes.NewBuffer(bd))
	// res, err := http.Post("http://localhost:3001/todos", "application/json", bytes.NewBuffer(bd))
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	resBD, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("response", string(resBD))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) getAllTodo() []string {
	res, err := http.Get("http://todolistbe:2344/todos")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer res.Body.Close()

	bd, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var todos []string
	err = json.Unmarshal(bd, &todos)
	if err != nil {
		log.Println(err)
		return nil
	}

	return todos
}
