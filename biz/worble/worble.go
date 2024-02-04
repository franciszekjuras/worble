package worble

import "slices"

const GuessCorrect = "correct"
const GuessPresent = "present"
const GuessIncorrect = "incorrect"

type Guess struct {
	Letter rune
	// should be one of Guess*
	Status string
}

type Answer [5]rune

type GuessScore [5]Guess

type Game struct {
	Guesses []GuessScore
}

// var validateGuess = regexp.MustCompile(`^[a-z]+$`).MatchString

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

func (game *Game) AddGuess(guessInput string) {
	guess := []rune(guessInput)
	if len(guess) != 5 {
		return
	}
	answer := Answer{'b', 'r', 'a', 'v', 'e'}
	guessScore := answer.scoreGuess(guess)
	game.Guesses = append(game.Guesses, guessScore)
}
