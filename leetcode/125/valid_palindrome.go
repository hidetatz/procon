package main

import (
	"strings"
	"unicode"
)

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	head, tail := 0, len(s)-1
	for {
		if head >= tail {
			break
		}

		for head < tail && !isAlphanumeric(rune(s[head])) {
			head++
		}

		for head < tail && !isAlphanumeric(rune(s[tail])) {
			tail--
		}

		if head < tail && strings.ToUpper(string(s[head])) != strings.ToUpper(string(s[tail])) {
			return false
		}
		head++
		tail--
	}

	return true

}
