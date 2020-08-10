package main

import "math"

func charToNum(s string) int {
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, c := range a {
		if string(c) == s {
			return i + 1
		}
	}

	panic("must not come here")
}

func titleToNumber(s string) int {
	l := len(s)

	ret := 0
	for i := 0; i < l; i++ {
		p := math.Pow(26, float64(i))
		ret += int(p) * charToNum(string(s[l-1-i]))
	}

	return ret
}
