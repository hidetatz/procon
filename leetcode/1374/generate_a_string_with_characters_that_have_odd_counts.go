package main

import "strings"

func generateTheString(n int) string {
	if n%2 != 0 {
		return strings.Repeat("a", n)
	}

	return strings.Repeat("a", n-1) + "b"
}
