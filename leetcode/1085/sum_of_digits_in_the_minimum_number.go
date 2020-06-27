package main

import (
	"fmt"
	"sort"
)

func sumOfDigits(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)
	min := A[0]

	fmt.Println(min)

	if min == 0 {
		return 0
	}

	sum := 0
	for min > 0 {
		mod := min % 10
		sum += mod
		min /= 10
	}

	if sum%2 == 0 {
		return 1
	}

	return 0
}
