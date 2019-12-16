package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type val struct {
	b, c int
}

type vals []val

func (v vals) Len() int           { return len(v) }
func (v vals) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v vals) Less(i, j int) bool { return v[i].c < v[j].c }

func main() {
	var sw = bufio.NewScanner(os.Stdin)
	sw.Split(bufio.ScanWords)

	sw.Scan()
	n, err := strconv.Atoi(sw.Text())
	check(err)

	sw.Scan()
	m, err := strconv.Atoi(sw.Text())
	check(err)

	cards := make([]int, n)

	for i := 0; i < n; i++ {
		sw.Scan()
		c, err := strconv.Atoi(sw.Text())
		check(err)
		cards[i] = c
	}

	var vs vals = make([]val, m)
	for i := 0; i < m; i++ {
		sw.Scan()
		b, err := strconv.Atoi(sw.Text())
		check(err)

		sw.Scan()
		c, err := strconv.Atoi(sw.Text())
		check(err)
		vs[i].b, vs[i].c = b, c
	}

	// ---

	sort.Sort(sort.Reverse(vs))
	cnt := 0
	for i := 0; i < m && cnt < n; i++ {
		for j := 0; j < vs[i].b && cnt < n; j++ {
			cards = append(cards, vs[i].c)
			cnt++
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	ans := 0
	for i := 0; i < n; i++ {
		ans += cards[i]
	}
	fmt.Println(ans)
}
