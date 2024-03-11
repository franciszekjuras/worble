package web

import (
	"net/http"

	"worble.ow6.foo/ui"
)

func (app *App) routes() http.Handler {
	// TODO: add session renewal middleware
	statefulMiddleware := app.SessionManager.LoadAndSave

	mux := http.NewServeMux()
	// stateless routes
	mux.HandleFunc("/", app.Home)

	// stateful routes
	mux.Handle("/play", statefulMiddleware(http.HandlerFunc(app.Play)))

	//static routes
	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("/static/", fileServer)
	return mux
}
