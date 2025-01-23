package main

import "net/http"

func (app *Application1) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", app.handleRegistration)
	mux.HandleFunc("/login", app.handleLogin)

	fileServer := http.FileServer(http.Dir("static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
