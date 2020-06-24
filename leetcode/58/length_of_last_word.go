package main

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			l++
			continue
		}

		break
	}

	return l
}
