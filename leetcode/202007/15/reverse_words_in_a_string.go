package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	s = regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(s), " ")
	ss := strings.Split(s, " ")

	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return strings.Join(ss, " ")
}
