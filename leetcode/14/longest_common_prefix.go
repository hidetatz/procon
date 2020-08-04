package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	i := 1
	prefix := ""
	for {
		temp := ""
		for _, str := range strs {
			if len(str) >= i {
				temp = str[:i]
				break
			}
		}

		if temp == "" {
			break
		}

		ok := true

		for _, str := range strs {
			if !strings.HasPrefix(str, temp) {
				ok = false
			}
		}

		if !ok {
			break
		}

		prefix = temp
		i++
		continue

	}

	return prefix

}
