package main

import (
	"bytes"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, td)
	if err != nil {
		http.Error(w, "internal server error: rendering", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
