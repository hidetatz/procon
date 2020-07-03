package main

import (
	"fmt"
)

func main() {
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
}

func prisonAfterNDays(cells []int, N int) []int {
	states := map[int]int{}
	states[concat(cells)] = 0

	found, lastFound := 0, 0

	for i := 1; i <= N; i++ {
		cells = nextDay(cells)

		key := concat(cells)
		if k, ok := states[key]; ok {
			found, lastFound = i, k
			break
		}

		states[key] = i
	}

	// when a state doesn't repeat
	if found-lastFound == 0 {
		return cells
	}

	mod := (N - found) % (found - lastFound)

	for i := 1; i <= mod; i++ {
		cells = nextDay(cells)
	}

	return cells
}

// transition the state of prison
func nextDay(prev []int) []int {
	now := make([]int, 8)
	for j := range prev {
		if j-1 < 0 || j+1 > 7 {
			now[j] = 0 // first/last cell
			continue
		}

		if (prev[j-1] == 0 && prev[j+1] == 0) || (prev[j-1] == 1 && prev[j+1] == 1) {
			now[j] = 1
		} else {
			now[j] = 0
		}
	}
	return now
}

// create integer by given array
// assuming length is always 8
func concat(arr []int) int {
	result := 0
	x := 10000000
	for i := 0; i < 8; i++ {
		result += arr[i] * x
		x /= 10
	}

	return result
}
