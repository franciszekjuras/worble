package main

import (
	"log"
	"net/http"

	"worble.ow6.foo/app/handlers"
)

func main() {
	err := handlers.InitTemplates()
	if err != nil {
		log.Fatal(err)
	}

	app := handlers.App{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/guess", app.PostGuess)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Starting server on :4001")
	err = http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
