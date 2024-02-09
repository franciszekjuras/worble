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

func (app *App) PlayGet(w http.ResponseWriter, r *http.Request) {
	err := app.Ts.ExecuteTemplate(w, "game-full.html", uimodels.MakeGame(&app.Game))
	if err != nil {
		log.Println(err.Error())
	}
}

func (app *App) PlayPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
	input := r.Form.Get("guess")
	app.Game.SubmitGuess(input)

	err = app.Ts.ExecuteTemplate(w, "game.html", uimodels.MakeGame(&app.Game))
	if err != nil {
		log.Println(err.Error())
	}

	// if app.Game.Result != nil {
	// 	app.Game = worble.NewGame()
	// }
}

func (app *App) PlayDelete(w http.ResponseWriter, r *http.Request) {
	app.Game = worble.NewGame()

	err := app.Ts.ExecuteTemplate(w, "game.html", uimodels.MakeGame(&app.Game))
	if err != nil {
		log.Println(err.Error())
	}
}

func (app *App) Play(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		app.PlayGet(w, r)
	} else if r.Method == http.MethodPost {
		app.PlayPost(w, r)
	} else if r.Method == http.MethodDelete {
		app.PlayDelete(w, r)
	}
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}
