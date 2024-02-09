package uimodels

var keyboardKeys = [][]rune{
	{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
	{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
	{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
}

type Keyboard struct {
	// one of worble.Guess* constants or empty
	Keys       [][]rune
	KeysStatus map[rune]int
}

func makeKeyboard(keysStatus map[rune]int) Keyboard {
	return Keyboard{Keys: keyboardKeys, KeysStatus: keysStatus}
}
