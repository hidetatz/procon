package main

func xorOperation(n int, start int) int {
	ret := start
	for i := 1; i < n; i++ {
		ret ^= (start + 2*i)
	}
	return ret
}
