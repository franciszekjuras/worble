package worble

const Rounds = 6
const WordLen = 5

const GuessCorrect = "correct"
const GuessPresent = "present"
const GuessIncorrect = "incorrect"

type Guess struct {
	Letter rune
	Points int
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

var letterScores = map[rune]int{
	'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'l': 1, 'n': 1, 's': 1, 't': 1, 'r': 1,
	'd': 2, 'g': 2, 'b': 3, 'c': 3, 'm': 3, 'p': 3,
	'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4, 'k': 5,
	'j': 8, 'x': 8, 'q': 10, 'z': 10,
}

func scorePresentLetter(letter rune) int {
	return letterScores[letter] / 2
}

func scoreCorrectLetter(letter rune) int {
	return letterScores[letter]
}

func guessSetPop(guessSet []rune, val rune) ([]rune, bool) {
	for i, letter := range guessSet {
		if letter == val {
			guessSet[i] = guessSet[len(guessSet)-1]
			return guessSet[:len(guessSet)-1], true
		}
	}
	return guessSet, false
}

func (answer *Answer) scoreGuess(guess []rune) GuessScore {
	score := GuessScore{}
	guessSet := make([]rune, 0, len(answer))
	for i, letter := range guess {
		score[i].Letter = letter
		if letter == answer[i] {
			score[i].Status = GuessCorrect
			score[i].Points = scoreCorrectLetter(letter)
		} else {
			guessSet = append(guessSet, answer[i])
		}
	}
	for i, letter := range guess {
		if score[i].Status == "" {
			var ok bool
			guessSet, ok = guessSetPop(guessSet, letter)
			if ok {
				score[i].Status = GuessPresent
				score[i].Points = scorePresentLetter(letter)
			} else {
				score[i].Status = GuessIncorrect
			}
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
