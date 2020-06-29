package main

import "strconv"

func selfDividingNumbers(left int, right int) []int {
	ret := []int{}

	for i := left; i <= right; i++ {
		if isSDN(i) {
			ret = append(ret, i)
		}
	}

	return ret
}

func isSDN(i int) bool {
	s := strconv.Itoa(i)
	for _, ss := range s {
		if ss == '0' {
			return false
		}

		si, _ := strconv.Atoi(string(ss))
		if i%si != 0 {
			return false
		}
	}

	return true
}
