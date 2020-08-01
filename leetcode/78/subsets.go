package main

import (
	"fmt"
)

func main() {
	// v := []int{9, 0, 3, 5, 7}
	// v := []int{1, 0, 3, 5, 7}
	v := []int{9, 0, 3, 5, 8}
	fmt.Println(subsets(v))
}

func subsets(nums []int) [][]int {
	result := [][]int{
		[]int{}, // empty
	}

	for _, num := range nums {
		temp := [][]int{}

		for _, r := range result {
			rr := make([]int, len(r))
			copy(rr, r)
			rr = append(rr, num)
			temp = append(temp, rr)
		}

		for _, t := range temp {
			result = append(result, t)
		}
	}

	return result
}
