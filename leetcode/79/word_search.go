package main

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if find(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func find(board [][]byte, row, col int, word string, index int) bool {
	if index >= len(word) {
		return true // every characters are already found
	}

	if row < 0 || col < 0 {
		return false
	}

	if row == len(board) || col == len(board[0]) {
		return false
	}

	if !(board[row][col] == word[index]) {
		return false
	}

	board[row][col] = '#'

	offsets := [][]int{
		{0, 1}, {1, 0}, {-1, 0}, {0, -1},
	}

	found := false
	for _, offset := range offsets {
		if find(board, row+offset[0], col+offset[1], word, index+1) {
			found = true
			break
		}
	}

	board[row][col] = word[index]
	return found
}
