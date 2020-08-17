package main

func restoreString(s string, indices []int) string {
	ret := make([]byte, len(indices))
	for i, index := range indices {
		ret[index] = s[i]
	}

	r := ""
	for _, rr := range ret {
		r += string(rr)
	}

	return r
}
