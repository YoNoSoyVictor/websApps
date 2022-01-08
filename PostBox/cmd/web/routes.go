package main

import "net/http"

func (app *app) routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/post", app.post)
	mux.HandleFunc("/post/create", app.create)

	mux.HandleFunc("/user/signup", app.signup)

	fs := http.FileServer(http.Dir("ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	return app.logRequest(secureHeaders(mux))
}
