package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			low, high := i+1, len(nums)-1
			for low < high {
				sum := nums[i] + nums[low] + nums[high]
				if sum < 0 || (low > i+1 && nums[low] == nums[low-1]) {
					low++
				} else if sum > 0 || (high < len(nums)-1 && nums[high] == nums[high+1]) {
					high--
				} else {
					result = append(result, []int{nums[i], nums[low], nums[high]})
					low++
					high--
				}
			}
		}
	}
	return result
}
