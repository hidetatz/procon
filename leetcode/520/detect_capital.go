package main

import (
	"fmt"
	"unicode"
)

func isUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func detectCapitalUse(word string) bool {
	if isUpper(string(word[0])) {
		return isUpper(string(word[1:])) || isLower(string(word[1:]))
	}

	return isLower(string(word[1:]))
}

func main() {
	cases := map[string]bool{
		"USA":      true,
		"leetcode": true,
		"Google":   true,
		"FlaG":     false,
	}

	for c, want := range cases {
		fmt.Printf("%v: want: %v, got: %v\n", c, want, detectCapitalUse(c))
	}
}
