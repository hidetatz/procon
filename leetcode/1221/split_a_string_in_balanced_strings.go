package main

func balancedStringSplit(s string) int {
	ret := 0
	cnt := 0

	for _, ss := range s {
		if ss == 'L' {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			ret++
		}
	}

	return ret
}
