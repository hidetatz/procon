package main

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}
