package web

import (
	"net/http"

	"worble.ow6.foo/ui"
)

func (app *App) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/play", app.Play)

	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("/static/", fileServer)
	return mux
}
