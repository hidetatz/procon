package main

func calculateTime(keyboard string, word string) int {
	m := map[rune]int{}

	for pos, key := range keyboard {
		m[key] = pos
	}

	now := 0

	ret := 0
	for _, c := range word {
		pos := m[c]
		move := pos - now
		if move < 0 {
			move = -move
		}

		ret += move
		now = pos
	}

	return ret
}
