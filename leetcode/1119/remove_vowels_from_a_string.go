package main

import (
	"fmt"
	"strings"
)

func removeVowels(S string) string {
	var b strings.Builder

	for _, ss := range S {
		if ss != 'a' &&
			ss != 'e' &&
			ss != 'i' &&
			ss != 'o' &&
			ss != 'u' {
			fmt.Fprintf(&b, string(ss))
		}
	}
	return b.String()

}
