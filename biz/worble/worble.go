package worble

import (
	"slices"
)

const Rounds = 6
const WordLen = 5

const GuessCorrect = "correct"
const GuessPresent = "present"
const GuessIncorrect = "incorrect"

type Guess struct {
	Letter rune
	// should be one of Guess*
	Status string
}

type Answer [WordLen]rune

type GuessScore [WordLen]Guess

type GameResult struct {
	FoundAnswer  bool
	NumOfGuesses int
}

type Game struct {
	Guesses []GuessScore
	Result  *GameResult
	Answer  Answer
}

func (guessScore *GuessScore) isComplete() bool {
	for i := 0; i < len(guessScore); i++ {
		if guessScore[i].Status != GuessCorrect {
			return false
		}
	}
	return true
}

func (answer *Answer) scoreGuess(guess []rune) GuessScore {
	score := GuessScore{}
	for i := 0; i < len(score); i++ {
		score[i].Letter = guess[i]
		if guess[i] == answer[i] {
			score[i].Status = GuessCorrect
		} else if slices.Contains(answer[:], guess[i]) {
			score[i].Status = GuessPresent
		} else {
			score[i].Status = GuessIncorrect
		}
	}
	return score
}

func NewGame() Game {
	return Game{Answer: Answer{'b', 'r', 'a', 'v', 'e'}}
}

func (game *Game) AddGuess(guessInput string) {
	if game.Result != nil {
		return
	}
	guess := []rune(guessInput)
	if len(guess) != WordLen {
		return
	}
	guessScore := game.Answer.scoreGuess(guess)
	game.Guesses = append(game.Guesses, guessScore)
	if guessScore.isComplete() {
		game.Result = &GameResult{FoundAnswer: true, NumOfGuesses: len(game.Guesses)}
	} else if len(game.Guesses) == Rounds {
		game.Result = &GameResult{FoundAnswer: false}
	}
}
