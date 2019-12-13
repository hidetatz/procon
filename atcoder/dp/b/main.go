package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	const inf = 1 << 60

	var sw = bufio.NewScanner(os.Stdin)
	sw.Split(bufio.ScanWords)
	sw.Scan()

	n, err := strconv.Atoi(sw.Text())
	if err != nil {
		panic(err)
	}

	sw.Scan()

	k, err := strconv.Atoi(sw.Text())
	if err != nil {
		panic(err)
	}

	h := make([]int, n)

	for i := 0; i < n; i++ {
		sw.Scan()
		num, err := strconv.Atoi(sw.Text())
		if err != nil {
			panic(err)
		}
		h[i] = num
	}

	var dp [100010]int
	for i := 0; i < 100010; i++ {
		dp[i] = inf
	}

	dp[0] = 0

	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				if dp[i] > dp[i-j]+abs(h[i]-h[i-j]) {
					dp[i] = dp[i-j] + abs(h[i]-h[i-j])
				}
			}
		}
	}

	fmt.Print(dp[n-1])
}
