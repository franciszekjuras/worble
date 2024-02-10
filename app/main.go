package main

import (
	"log"
	"net/http"

	"worble.ow6.foo/app/handlers"
	"worble.ow6.foo/appui/uitempl"
	"worble.ow6.foo/biz/worble"
	"worble.ow6.foo/ui"
)

func main() {
	ts, err := uitempl.InitTemplates()
	if err != nil {
		log.Fatalln("Failed to initialize templates: ", err)
	}

	app := handlers.App{Ts: ts, Game: worble.NewGame()}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/play", app.Play)

	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("/static/", fileServer)

	log.Println("Starting server on :4001")
	err = http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
