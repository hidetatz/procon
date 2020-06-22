package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := range haystack {
		if len(haystack[i:])-len(needle) < 0 {
			return -1
		}

		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
