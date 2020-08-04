package main

import "math"

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}

	answers := map[int]struct{}{}
	for i := 0; ; i++ {
		p := math.Pow(4, float64(i))
		if p > math.MaxInt32 {
			break
		}
		answers[int(p)] = struct{}{}
	}

	_, ok := answers[num]
	return ok
}
