package worble

const GuessCorrect = "correct"
const GuessPresent = "present"
const GuessIncorrect = "incorrect"

type Guess struct {
	Letter rune
	// should be one of Guess*
	Status string
}

type GuessWord [5]Guess

type Game struct {
	Guesses []GuessWord
}
