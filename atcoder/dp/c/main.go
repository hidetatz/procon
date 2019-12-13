package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sw = bufio.NewScanner(os.Stdin)
	sw.Split(bufio.ScanWords)
	sw.Scan()

	n, err := strconv.Atoi(sw.Text())
	if err != nil {
		panic(err)
	}

	sv := make([][]int, n)

	for i := 0; i < n; i++ {
		sw.Scan()
		a, err := strconv.Atoi(sw.Text())
		if err != nil {
			panic(err)
		}

		sw.Scan()
		b, err := strconv.Atoi(sw.Text())
		if err != nil {
			panic(err)
		}

		sw.Scan()
		c, err := strconv.Atoi(sw.Text())
		if err != nil {
			panic(err)
		}
		sv[i] = []int{a, b, c}
	}

	// dp[i][0] = maximum happiness on date `i` when selected a
	// dp[i][1] = maximum happiness on date `i` when selected b
	// dp[i][2] = maximum happiness on date `i` when selected c
	var dp [100010][3]int
	for i := 0; i < 100010; i++ {
		dp[i] = [3]int{}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j != k {
					if dp[i+1][k] < dp[i][j]+sv[i][k] {
						dp[i+1][k] = dp[i][j] + sv[i][k]
					}
				}
			}
		}
	}

	res := 0
	if res < dp[n][0] {
		res = dp[n][0]
	}

	if res < dp[n][1] {
		res = dp[n][1]
	}

	if res < dp[n][2] {
		res = dp[n][2]
	}

	fmt.Print(res)
}
