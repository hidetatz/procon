package main

func shuffle(nums []int, n int) []int {
	ret := []int{}
	for i := 0; i < n; i++ {
		ret = append(ret, nums[i], nums[i+n])
	}

	return ret
}
