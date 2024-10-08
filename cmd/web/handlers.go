package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.elmm.net/internal/models"
)

func home(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		snippets, err := app.snippets.Latest()
		if err != nil {
			app.ServerError(w, r, err)
			return
		}
		w.Write([]byte(fmt.Sprintf("%v", snippets)))
	}
}

func snippetView(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		sp, err := app.snippets.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				http.NotFound(w, r)
			} else {
				app.ServerError(w, r, err)
			}
			return
		}
		w.Write([]byte(fmt.Sprintf("Display a specific snippet with %v is %v ...", id, sp)))
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
		title, content, expires := "0 snail", "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa", 7

		id, err := app.snippets.Insert(title, content, expires)
		if err != nil {
			app.ServerError(w, r, err)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	}
}
