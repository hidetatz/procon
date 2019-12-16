package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var sw = bufio.NewScanner(os.Stdin)
	sw.Split(bufio.ScanWords)

	sw.Scan()
	n, err := strconv.Atoi(sw.Text())
	check(err)

	sw.Scan()
	m, err := strconv.Atoi(sw.Text())
	check(err)

	ls := make([]int, m)
	rs := make([]int, m)

	for i := 0; i < m; i++ {
		sw.Scan()
		l, err := strconv.Atoi(sw.Text())
		check(err)
		ls[i] = l

		sw.Scan()
		r, err := strconv.Atoi(sw.Text())
		check(err)
		rs[i] = r
	}

	// ---

	left, right := 0, n
	for i := 0; i < m; i++ {
		l, r := ls[i], rs[i]
		if l > left {
			left = l
		}

		if right > r {
			right = r
		}
	}
	if left <= right {
		fmt.Print(right - left + 1)
		return
	}

	fmt.Print(0)
}
