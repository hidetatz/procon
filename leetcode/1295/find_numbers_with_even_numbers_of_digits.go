package main

func findNumbers(nums []int) int {
	ret := 0

	for _, num := range nums {
		if (10 <= num && num <= 99) || (1000 <= num && num <= 9999) || num == 100000 {
			ret++
		}
	}

	return ret
}
