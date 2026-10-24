package main

import "net/http"

// The routes() method returns a servermux containing our application routes
func (app *application) routes() *http.ServeMux {
   	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir(cfg.staticDir)})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/gist/view", app.gistView)
	mux.HandleFunc("/gist/create", app.gistCreate)

    return mux
}
