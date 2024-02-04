package handlers

import (
	"html/template"
	"log"
	"net/http"

	"worble.ow6.foo/appui/uimodels"
	"worble.ow6.foo/biz/worble"
)

type App struct {
	Game worble.Game
	Ts   *template.Template
}

func (app *App) PostGuess(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
	input := r.Form.Get("guess")
	log.Println("input:", input)
	app.Game.AddGuess(input)

	err = app.Ts.ExecuteTemplate(w, "game.html", uimodels.MakeBoard(app.Game))
	if err != nil {
		log.Println(err.Error())
	}
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := app.Ts.ExecuteTemplate(w, "game-full.html", uimodels.MakeBoard(app.Game))
	if err != nil {
		log.Println(err.Error())
	}
}
