package uimodels

import "worble.ow6.foo/biz/worble"

type BoardRowFilled struct {
	Guess         worble.GuessScore
	Bonus         []int
	Animate       bool
	AutoSubmitted bool
}

type BoardRowInput struct {
	Bonus []int
}

type BoardRowEmpty struct {
}

type Board [worble.Rounds]map[string]any

func makeBoard(game *worble.Game) Board {
	var board Board
	guessNum := len(game.Guesses)
	autoSubmittedIdx := guessNum
	if game.Result != nil && game.Result.FoundAnswer {
		autoSubmittedIdx = game.Result.NumOfGuesses
	}
	for i, guess := range game.Guesses {
		board[i] = map[string]any{
			"Filled": BoardRowFilled{
				Guess: guess, Bonus: game.Bonuses[i][:], Animate: i == guessNum-1, AutoSubmitted: i >= autoSubmittedIdx,
			},
		}
	}
	emptyRowsStart := guessNum
	if guessNum < worble.Rounds && game.Result == nil {
		board[guessNum] = map[string]any{"Input": BoardRowInput{Bonus: game.Bonuses[guessNum][:]}}
		emptyRowsStart++
	}
	for i := emptyRowsStart; i < worble.Rounds; i++ {
		board[i] = map[string]any{"Empty": BoardRowEmpty{}}
	}
	return board
}
