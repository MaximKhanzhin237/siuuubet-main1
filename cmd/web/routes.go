package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", app.home)
	mux.HandleFunc("/", app.get)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
