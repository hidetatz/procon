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
	a, err := strconv.Atoi(sw.Text())
	check(err)

	sw.Scan()
	b, err := strconv.Atoi(sw.Text())
	check(err)

	// ---

	if 13 <= a {
		fmt.Print(b)
		return
	}

	if 6 <= a && a <= 12 {
		fmt.Print(b / 2)
		return
	}

	fmt.Print(0)
}
