package main

func singleNumber(nums []int) []int {
	m := map[int]int{}

	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num] = 2
		}
	}

	ret := []int{}

	for num, found := range m {
		if found == 1 {
			ret = append(ret, num)
			if len(ret) == 2 {
				break
			}
		}
	}

	return ret
}
