package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	half := (len(nums) + 1) / 2

	if len(nums)%2 != 0 {
		return float64(nums[half-1])
	}

	return float64((nums[half-1] + nums[half])) / 2
}

func main() {
	type s struct {
		nums1 []int
		nums2 []int
	}
	tc := []s{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
		},
	}
	for _, t := range tc {
		n1 := append([]int{}, t.nums1...)
		n2 := append([]int{}, t.nums2...)
		fmt.Printf("nums1: %v, nums2: %v => %v\n", n1, n2, findMedianSortedArrays(t.nums1, t.nums2))
	}
}
