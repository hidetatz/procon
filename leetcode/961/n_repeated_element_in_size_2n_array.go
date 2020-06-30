package main

func repeatedNTimes(A []int) int {
	m := map[int]int{}

	for _, e := range A {
		if _, ok := m[e]; ok {
			m[e]++
		} else {
			m[e] = 1
		}
	}

	for k, v := range m {
		if v == len(A)/2 {
			return k
		}
	}

	return 0
}
