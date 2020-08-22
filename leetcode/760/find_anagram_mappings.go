package main

func anagramMappings(A []int, B []int) []int {
	m := map[int]int{}

	for i, bb := range B {
		m[bb] = i
	}

	ret := make([]int, len(A))
	for i, aa := range A {
		ret[i] = m[aa]
	}

	return ret
}
