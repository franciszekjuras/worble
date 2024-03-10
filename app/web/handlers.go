package web

import (
	"encoding/gob"
	"html/template"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"worble.ow6.foo/appui/uimodels"
	"worble.ow6.foo/biz/worble"
)

type App struct {
	Ts             *template.Template
	SessionManager *scs.SessionManager
}

func init() {
	gob.Register(worble.Game{})
}

func (app *App) PlayGet(w http.ResponseWriter, r *http.Request) {
	sm := app.SessionManager
	ctx := r.Context()

	var game worble.Game
	if !sm.Exists(ctx, "game") {
		game = worble.NewGame()
		sm.Put(ctx, "game", game)
	} else {
		game = sm.Get(ctx, "game").(worble.Game)
	}
	err := app.Ts.ExecuteTemplate(w, "game-full.html", uimodels.MakeGame(&game))
	if err != nil {
		log.Println(err.Error())
	}
}

func (app *App) PlayPost(w http.ResponseWriter, r *http.Request) {
	sm := app.SessionManager
	ctx := r.Context()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}

	game := sm.Get(ctx, "game").(worble.Game) //TODO: validate

	input := r.Form.Get("guess")
	game.SubmitGuess(input)

	sm.Put(ctx, "game", game)

	err = app.Ts.ExecuteTemplate(w, "game.html", uimodels.MakeGame(&game))
	if err != nil {
		log.Println(err.Error())
	}
}

func (app *App) PlayDelete(w http.ResponseWriter, r *http.Request) {
	sm := app.SessionManager
	ctx := r.Context()

	game := worble.NewGame()
	sm.Put(ctx, "game", game)

	err := app.Ts.ExecuteTemplate(w, "game.html", uimodels.MakeGame(&game))
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
