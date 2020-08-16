package main

func numJewelsInStones(J string, S string) int {
	m := map[rune]struct{}{}
	for _, jj := range J {
		m[jj] = struct{}{}
	}

	ret := 0
	for _, ss := range S {
		if _, ok := m[ss]; ok {
			ret++
		}
	}

	return ret
}
