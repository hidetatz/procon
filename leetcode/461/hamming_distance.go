package main

import "math/bits"

func hammingDistance(x int, y int) int {
	return int(bits.OnesCount(uint(x ^ y)))
}
