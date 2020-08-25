package main

func numIdenticalPairs(nums []int) int {
	ret := 0
	m := map[int]int{}
	for _, num := range nums {
		n, ok := m[num]
		if ok {
			ret += n
			m[num]++
		} else {
			m[num] = 1
		}
	}

	return ret
}
