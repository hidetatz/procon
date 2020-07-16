package main

import "math"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if x == 1 {
		return 1
	}

	if x == 0 {
		if n > 0 {
			return 0
		}

		if n < 0 {
			return math.Inf(int(x))
		}

		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	var ans float64 = 1
	var cur float64 = x

	for i := n; i > 0; i /= 2 {
		if (i % 2) == 1 {
			ans = ans * cur
		}

		cur *= cur
	}

	return ans
}
