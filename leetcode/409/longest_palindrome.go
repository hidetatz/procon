package main

func longestPalindrome(s string) int {
	m := map[string]int{}
	for _, ss := range s {
		if _, ok := m[string(ss)]; ok {
			m[string(ss)]++
		} else {
			m[string(ss)] = 1
		}
	}

	ret := 0
	for c, cnt := range m {
		if cnt%2 == 0 {
			ret += cnt
			delete(m, c)
		} else if cnt >= 3 {
			ret += cnt - 1
			m[c] = 1
		}
	}

	if len(m) > 0 {
		ret++
	}

	return ret
}
