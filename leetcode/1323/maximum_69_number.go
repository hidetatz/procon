package main

import (
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	r, _ := strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return r
}
