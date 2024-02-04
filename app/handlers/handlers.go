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
		board[i] = map[string]any{"Filled": BoardRowFilled{guess, i == guessNum-1}}
	}
	if guessNum < 5 {
		board[guessNum] = map[string]any{"Input": BoardRowInput{}}
	}
	for i := guessNum + 1; i < 5; i++ {
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

	// var game worble.Game
	// newGuess := worble.GuessWord{{'g', worble.GuessCorrect}, {'u', worble.GuessIncorrect},
	// 	{'e', worble.GuessPresent}, {'s', worble.GuessIncorrect}, {'s', worble.GuessIncorrect}}
	// game.Guesses = append(game.Guesses, newGuess)
	err := ts.ExecuteTemplate(w, "game-full.html", makeBoard(app.game))
	if err != nil {
		log.Println(err.Error())
	}
}
