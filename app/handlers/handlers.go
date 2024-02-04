package handlers

import (
	"html/template"
	"log"
	"net/http"

	"worble.ow6.foo/biz/worble"
	"worble.ow6.foo/com/templatefuncs"
)

type App struct {
	game worble.Game
}

var ts *template.Template

func InitTemplates() error {
	funcMap := template.FuncMap{
		"sub":  templatefuncs.Sub,
		"span": templatefuncs.Span,
	}
	ts_, err := template.New("").Funcs(funcMap).ParseGlob("./ui/html/*.html")
	ts = ts_
	return err
}

type BoardRowFilled struct {
	Guess   worble.GuessScore
	Animate bool
}

type BoardRowInput struct {
}

type BoardRowEmpty struct{}

type Board [5]map[string]any

func makeBoard(game worble.Game) Board {
	var board Board
	guessNum := len(game.Guesses)
	for i, guess := range game.Guesses {
		board[i] = map[string]any{"Filled": BoardRowFilled{Guess: guess, Animate: i == guessNum-1}}
	}
	emptyRowsStart := guessNum
	if guessNum < 5 && game.Result == nil {
		board[guessNum] = map[string]any{"Input": BoardRowInput{}}
		emptyRowsStart++
	}
	for i := emptyRowsStart; i < 5; i++ {
		board[i] = map[string]any{"Empty": BoardRowEmpty{}}
	}
	return board
}

func (app *App) PostGuess(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
	}
	input := r.Form.Get("guess")
	log.Println("input:", input)
	app.game.AddGuess(input)

	err = ts.ExecuteTemplate(w, "game.html", makeBoard(app.game))
	if err != nil {
		log.Println(err.Error())
	}
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := ts.ExecuteTemplate(w, "game-full.html", makeBoard(app.game))
	if err != nil {
		log.Println(err.Error())
	}
}
