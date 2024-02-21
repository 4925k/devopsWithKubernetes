package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	todos := app.getAllTodo()

	app.render(w, r, "home.page.tmpl", &templateData{Todos: todos})
}
