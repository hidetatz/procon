package main

import "math"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		last := x % 10
		x /= 10

		if ret > math.MaxInt32/10 || ret < math.MinInt32/10 {
			return 0
		}

		ret = ret*10 + last
	}

	return ret
}
