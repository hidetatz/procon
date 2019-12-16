package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func chMax(i *int, b int) {
	if *i < b {
		*i = b
	}
}

// func chMin(i *int, b int) {
// 	if *i > b {
// 		*i = b
// 	}
// }

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
	w, err := strconv.Atoi(sw.Text())
	check(err)

	ws := make([]int, n)
	vs := make([]int, n)

	for i := 0; i < n; i++ {
		sw.Scan()
		w, err := strconv.Atoi(sw.Text())
		check(err)
		ws[i] = w

		sw.Scan()
		v, err := strconv.Atoi(sw.Text())
		check(err)
		vs[i] = v
	}

	// ---

	// dp[i][j] => Maximum value when picking up 0 to i items and value must be under j
	//
	// * i = (0, 1, ... n)
	// * j = (0, 1, ... w)
	var dp [110][100100]int

	for i := 0; i < n; i++ {
		for j := 1; j <= w; j++ {
			if j-ws[i] >= 0 {
				// When picking up i+1 item
				chMax(&dp[i+1][j], dp[i][j-ws[i]]+vs[i])
			}

			// When not
			chMax(&dp[i+1][j], dp[i][j])
		}
	}
	fmt.Print(dp[n][w])
}
