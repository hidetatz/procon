package main

func judgeCircle(moves string) bool {
	pos := []int{0, 0}
	for _, move := range moves {
		switch move {
		case 'U':
			pos[0]++
		case 'D':
			pos[0]--
		case 'L':
			pos[1]--
		case 'R':
			pos[1]++
		}
	}

	return pos[0] == 0 && pos[1] == 0
}
