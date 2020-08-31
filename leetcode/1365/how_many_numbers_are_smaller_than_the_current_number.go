package main

// O(n^2)
func smallerNumbersThanCurrent(nums []int) []int {
	ret := []int{}
	for _, num := range nums {
		cnt := 0
		for _, num2 := range nums {
			if num2 < num {
				cnt++
			}
		}

		ret = append(ret, cnt)
	}

	return ret
}
