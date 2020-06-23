package main

func searchInsert(nums []int, target int) int {
	idx := 0
	found := false
	for i, num := range nums {
		if target == num {
			return i
		}

		if num < target {
			idx = i
			found = true
		}
	}

	if !found {
		return 0
	}

	return idx + 1
}
