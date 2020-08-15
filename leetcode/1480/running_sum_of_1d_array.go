package main

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	ret := []int{nums[0]}

	for i := 1; i < len(nums)-1; i++ {
		ret = append(ret, nums[i]+ret[i-1])
	}

	ret = append(ret, nums[len(nums)-1]+ret[len(ret)-1])

	return ret
}
