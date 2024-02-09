package uimodels

import "worble.ow6.foo/biz/worble"

func MapGuessCodeToClass(code int) string {
	switch code {
	case worble.GuessIncorrect:
		return "incorrect"
	case worble.GuessPresent:
		return "present"
	case worble.GuessCorrect:
		return "correct"
	default:
		return ""
	}
}
