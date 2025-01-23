package main

import "net/http"

func (app *Application2) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/inf", app.inf)

	fileServer := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
