package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {
	uglyNums := make([]int, n)

	i2, i3, i5 := 0, 0, 0

	nextMultiple2 := 2
	nextMultiple3 := 3
	nextMultiple5 := 5

	nextUglyNo := 1

	uglyNums[0] = 1

	for i := 1; i < n; i++ {
		nextUglyNo = min(nextMultiple2, min(nextMultiple3, nextMultiple5))

		uglyNums[i] = nextUglyNo
		if nextUglyNo == nextMultiple2 {
			i2 += 1
			nextMultiple2 = uglyNums[i2] * 2
		}

		if nextUglyNo == nextMultiple3 {
			i3 += 1
			nextMultiple3 = uglyNums[i3] * 3
		}

		if nextUglyNo == nextMultiple5 {
			i5 += 1
			nextMultiple5 = uglyNums[i5] * 5
		}
	}

	fmt.Println(uglyNums)

	return nextUglyNo
}

func min(n, m int) int {
	if n < m {
		return n
	}

	return m
}
