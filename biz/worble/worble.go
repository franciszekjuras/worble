package worble

import "math/rand"

const Rounds = 6
const WordLen = 5

const BonusNone = 1
const BonusDouble = 2
const BonusTriple = 3

const GuessIncorrect = 1
const GuessPresent = 2
const GuessCorrect = 3

type Guess struct {
	Letter rune
	Points int
	// should be one of Guess*
	Status int
}

type Answer [WordLen]rune

type GuessScore [WordLen]Guess

type GameResult struct {
	FoundAnswer  bool
	NumOfGuesses int
	Points       int
}

type Bonuses [Rounds][WordLen]int

type Game struct {
	Guesses       []GuessScore
	Result        *GameResult
	LettersStatus map[rune]int
	Bonuses       *Bonuses
	Answer        Answer
}

var bonusTable = Bonuses{{2, 1, 3, 1, 2}, {1, 3, 1, 2, 1}, {2, 1, 1, 3, 1}, {1, 1, 2, 1, 2}, {1, 2, 1, 1, 1}, {1, 1, 1, 1, 1}}

func NewGame() Game {
	game := Game{LettersStatus: map[rune]int{}, Bonuses: &bonusTable}
	answer := wordsAnswers[rand.Intn(len(wordsAnswers))]
	copy(game.Answer[:], []rune(answer))
	return game
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

func scorePresentLetter(letter rune, bonus int) int {
	return bonus * letterScores[letter] / 2
}

func scoreCorrectLetter(letter rune, bonus int) int {
	return bonus * letterScores[letter]
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

func (game *Game) scoreGuess(guess []rune) GuessScore {
	bonus := game.Bonuses[len(game.Guesses)]
	answer := &game.Answer
	score := GuessScore{}
	guessSet := make([]rune, 0, len(answer))
	for i, letter := range guess {
		score[i].Letter = letter
		if letter == answer[i] {
			score[i].Status = GuessCorrect
			score[i].Points = scoreCorrectLetter(letter, bonus[i])
		} else {
			guessSet = append(guessSet, answer[i])
		}
	}
	for i, letter := range guess {
		if score[i].Status == 0 {
			var ok bool
			guessSet, ok = guessSetPop(guessSet, letter)
			if ok {
				score[i].Status = GuessPresent
				score[i].Points = scorePresentLetter(letter, bonus[i])
			} else {
				score[i].Status = GuessIncorrect
			}
		}
	}
	return score
}

func (game *Game) updateLettersStatus(score *GuessScore) {
	for i := range score {
		s := &score[i]
		game.LettersStatus[s.Letter] = max(s.Status, game.LettersStatus[s.Letter])
	}
}

func (game *Game) totalScore() int {
	score := 0
	for _, guess := range game.Guesses {
		for _, letter := range guess {
			score += letter.Points
		}
	}
	return score
}

func (game *Game) SubmitGuess(guessInput string) {
	if game.Result != nil {
		return
	}
	guess := []rune(guessInput)
	if len(guess) != WordLen {
		return
	}
	if _, ok := wordsValid[guessInput]; !ok {
		return
	}
	guessScore := game.scoreGuess(guess)
	game.updateLettersStatus(&guessScore)
	game.Guesses = append(game.Guesses, guessScore)
	if guessScore.isComplete() {
		numofGuesses := len(game.Guesses)
		for i := len(game.Guesses); i < Rounds; i++ {
			guessScore := game.scoreGuess(guess)
			game.Guesses = append(game.Guesses, guessScore)
		}
		game.Result = &GameResult{FoundAnswer: true, NumOfGuesses: numofGuesses, Points: game.totalScore()}
	} else if len(game.Guesses) == Rounds {
		game.Result = &GameResult{FoundAnswer: false}
	}
}
