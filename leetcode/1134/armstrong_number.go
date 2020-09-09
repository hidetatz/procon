package main

import "math"

func isArmstrong(N int) bool {
	orig := N
	digits := 0
	nums := []int{}

	for N != 0 {
		nums = append(nums, N%10)
		N /= 10
		digits++
	}

	result := 0
	for _, num := range nums {
		s := int(math.Pow(float64(num), float64(digits)))
		result += s
	}

	return result == orig
}
