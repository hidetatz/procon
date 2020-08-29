package main

func decompressRLElist(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	ret := []int{}
	for i := 0; i < len(nums); i++ {
		freq := nums[i]
		i++
		val := nums[i]
		for j := 0; j < freq; j++ {
			ret = append(ret, val)
		}
	}

	return ret
}
