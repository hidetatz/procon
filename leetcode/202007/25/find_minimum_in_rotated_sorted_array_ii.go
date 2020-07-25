package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMin([]int{3, 1, 3}))
}

func findMin2(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func findMin(nums []int) int {
	uniq := []int{nums[0]}
	for _, num := range nums[1:] {
		if uniq[len(uniq)-1] != num {
			uniq = append(uniq, num)
		}
	}

	if len(uniq) == 1 {
		return uniq[0]
	}

	// for [1, 2, 3, 1]
	if uniq[0] == uniq[len(uniq)-1] {
		uniq = uniq[1:]
	}

	// for [1, 1]
	if uniq[0] < uniq[len(uniq)-1] {
		return uniq[0]
	}

	return binarySearch(uniq, 0, len(uniq)-1)

}

func binarySearch(arr []int, l, r int) int {
	mid := l + (r-l)/2

	// e.g. [3, 4, 5, 1, 2]
	// mid = 2
	if arr[mid] > arr[mid+1] {
		return arr[mid+1]
	}

	// e.g. [4, 5, 1, 2, 3]
	// mid = 2
	if arr[mid-1] > arr[mid] {
		return arr[mid]
	}

	// e.g. [2, 3, 4, 5, 1]
	// mid = 2
	if arr[mid] > arr[0] {
		return binarySearch(arr, mid+1, r)
	}

	// e.g. [5, 1, 2, 3, 4]
	// mid = 2
	return binarySearch(arr, l, mid-1)
}
