package main

import (
	"fmt"
	"strings"
)

// without memo, time limit exceeded happens when
// s:    "aaaaaaaaaaaaaaaaaaaaaaa..."
// dict: ["a","aa","aaa","aaaa","aaaaa",...]
var memo map[string][]string

func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 {
		return []string{}
	}

	memo = map[string][]string{}
	return helper(s, wordDict)
}

func helper(s string, dict []string) []string {
	if m, ok := memo[s]; ok {
		return m
	}

	if s == "" {
		return nil
	}

	ret := []string{}

	for _, word := range dict {
		// when s == "catsanddog", dict = ["cat", "cats", "and", "sand", "dog"]
		// first, find leading "cat"
		if strings.HasPrefix(s, word) {
			// then, call this function again with s == "sanddog"
			rest := helper(strings.Replace(s, word, "", 1), dict)

			// nil means the "word" is not found in s
			if rest == nil {
				ret = append(ret, word)
				continue
			}

			// after recursion, rest should be like ["sand dog"]
			for _, w := range rest {
				// ret will be "cat sand dog"
				ret = append(ret, fmt.Sprintf("%s %s", word, w))
			}
		}
	}
	memo[s] = ret
	return ret
}
