package main

func removeDuplicates(nums []int) int {
	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}
