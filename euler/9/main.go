package main

import "fmt"

func main() {
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			for c := 0; c < 1000; c++ {
				if !(a < b && b < c) {
					continue
				}

				if a+b+c != 1000 {
					continue
				}

				if (a*a)+(b*b) != (c * c) {
					continue
				}
				fmt.Println(a * b * c)
			}
		}
	}
}
