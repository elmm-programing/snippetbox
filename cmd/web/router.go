package main

import "net/http"

func (app *Application) router() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home(app))
	mux.HandleFunc("GET /snippet/view/{id}", snippetView(app))
	mux.HandleFunc("GET /snippet/create", snippetCreate(app))
	mux.HandleFunc("POST /snippet/create", snippetCreatePost(app))
	return mux
}
