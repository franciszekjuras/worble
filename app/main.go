package main

import (
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on", port)
	err = http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
