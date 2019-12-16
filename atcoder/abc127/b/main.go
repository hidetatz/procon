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
	r, err := strconv.Atoi(sw.Text())
	check(err)

	sw.Scan()
	d, err := strconv.Atoi(sw.Text())
	check(err)

	sw.Scan()
	x, err := strconv.Atoi(sw.Text())
	check(err)

	// ---

	previous := x
	for i := 0; i < 10; i++ {
		ans := r*previous - d
		fmt.Println(ans)
		previous = ans
	}
}
