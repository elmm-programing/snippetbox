package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

)

func home(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "go")
		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/pages/home.tmpl.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
      app.ServerError(w, r, err )
			return

		}
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
      app.ServerError(w, r, err )
		}
	}
}

func snippetView(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte(fmt.Sprintf("Display a specific snippet with %v ...", id)))
	}
}

func snippetCreate(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Display a form for creating a new snippet..."))
	}
}

func typeAssertion[T string | int](value T) string {
	var val interface{} = value
	if _, ok := val.(string); ok {
		return "string"
	} else {
		return "int"
	}
}

func snippetCreatePost(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Save new snippet"))
	}
}
