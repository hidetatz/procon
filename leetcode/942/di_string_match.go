package main

func diStringMatch(S string) []int {
	ret := []int{}
	min, max := 0, len(S)

	for _, ss := range S {
		if ss == 'I' {
			ret = append(ret, min)
			min++
		} else {
			ret = append(ret, max)
			max--
		}
	}

	ret = append(ret, min)

	return ret
}
