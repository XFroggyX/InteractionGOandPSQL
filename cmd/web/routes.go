package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/table", app.showTable)
	mux.HandleFunc("/table/insert", app.insertTable)

	fileServer := http.FileServer(customizableFileSystem{http.Dir("./static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
