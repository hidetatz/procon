package main

func findDuplicates(nums []int) []int {
	ret := []int{}
	m := map[int]int{}
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 0
		} else {
			m[n]++
			ret = append(ret, n)
		}
	}
	return ret
}
