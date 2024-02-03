package handlers

import (
	"html/template"
	"log"
	"net/http"

	"worble.ow6.foo/biz/worble"
	"worble.ow6.foo/com/templatefuncs"
)

// type App struct {
// 	Game worble.Game
// }

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

func PostGuess(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	} else {
		input := r.Form.Get("guess")
		log.Println(input)
	}
	http.NotFound(w, r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	var game worble.Game
	// newGuess := worble.GuessWord{{'g', worble.GuessCorrect}, {'u', worble.GuessIncorrect},
	// 	{'e', worble.GuessPresent}, {'s', worble.GuessIncorrect}, {'s', worble.GuessIncorrect}}
	// game.Guesses = append(game.Guesses, newGuess)
	err := ts.ExecuteTemplate(w, "game-full.html", game)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
