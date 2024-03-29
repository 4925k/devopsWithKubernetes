package main

import (
	"html/template"
	"path/filepath"
	"time"
)

type templateData struct {
	Todos []string
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

// humanDate returns the given time in a human readable string format
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		// if err != nil {
		// 	return nil, err
		// }

		cache[name] = ts
	}

	return cache, nil
}
