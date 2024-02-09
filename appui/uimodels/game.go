package uimodels

import "worble.ow6.foo/biz/worble"

type Game struct {
	Board    Board
	Keyboard Keyboard
	Result   *worble.GameResult
}

func MakeGame(game *worble.Game) Game {
	var gameUi Game
	gameUi.Board = makeBoard(game)
	gameUi.Result = game.Result
	gameUi.Keyboard = makeKeyboard(game.LettersStatus)
	return gameUi
}
