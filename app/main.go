package main

import (
	"log"
	"net/http"

	"worble.ow6.foo/app/handlers"
	"worble.ow6.foo/appui/uitempl"
)

func main() {
	ts, err := uitempl.InitTemplates()
	if err != nil {
		log.Fatalln("Failed to initialize templates: ", err)
	}

	app := handlers.App{Ts: ts}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/guess", app.PostGuess)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Starting server on :4001")
	err = http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
