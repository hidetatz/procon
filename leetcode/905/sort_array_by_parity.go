package main

func sortArrayByParity(A []int) []int {
	ret := make([]int, len(A))
	f, l := 0, len(A)-1
	for _, aa := range A {
		if aa%2 == 0 {
			ret[f] = aa
			f++
		} else {
			ret[l] = aa
			l--
		}
	}

	return ret
}
