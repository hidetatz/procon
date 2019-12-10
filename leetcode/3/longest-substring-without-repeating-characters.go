package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	ans := 0
	carry := map[string]int{}
	ss := strings.Split(s, "")
	i, j := 0, 0
	for i < len(ss) && j < len(ss) {
		if _, ok := carry[ss[j]]; !ok {
			carry[ss[j]] = 0 // no meaning
			j++
		} else {
			delete(carry, ss[i])
			i++
		}
		if ans < len(carry) {
			ans = len(carry)
		}
	}
	return ans
}

func main() {
	tc := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}

	for _, t := range tc {
		fmt.Printf("%s = %+v\n", t, lengthOfLongestSubstring(t))
	}
}
