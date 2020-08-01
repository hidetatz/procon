package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	result := ""
	s := 0

	al := len(a) - 1
	bl := len(b) - 1

	for al >= 0 || bl >= 0 || s == 1 {
		if al >= 0 {
			aa, _ := strconv.Atoi(string(a[al]))
			s += aa
		}

		if bl >= 0 {
			bb, _ := strconv.Atoi(string(b[bl]))
			s += bb
		}

		result = fmt.Sprintf("%d%s", s%2, result)

		s /= 2

		al--
		bl--
	}

	return result
}
