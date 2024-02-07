package uimodels

import "worble.ow6.foo/biz/worble"

type BoardRowFilled struct {
	Guess   worble.GuessScore
	Animate bool
}

type BoardRowInput struct {
}

type BoardRowEmpty struct{}

type Board [worble.Rounds]map[string]any

type Game struct {
	Board  Board
	Result *worble.GameResult
}

func makeBoard(game *worble.Game) Board {
	var board Board
	guessNum := len(game.Guesses)
	for i, guess := range game.Guesses {
		board[i] = map[string]any{"Filled": BoardRowFilled{Guess: guess, Animate: i == guessNum-1}}
	}
	emptyRowsStart := guessNum
	if guessNum < worble.Rounds && game.Result == nil {
		board[guessNum] = map[string]any{"Input": BoardRowInput{}}
		emptyRowsStart++
	}
	for i := emptyRowsStart; i < worble.Rounds; i++ {
		board[i] = map[string]any{"Empty": BoardRowEmpty{}}
	}
	return board
}

func MakeGame(game *worble.Game) Game {
	var gameUi Game
	gameUi.Board = makeBoard(game)
	gameUi.Result = game.Result
	return gameUi
}
