package main

import (
	"fmt"
	"math"
)

func main() {
	primes := []int{}

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		} else if n == 2 {
			return true
		} else if n%2 == 0 {
			return false
		}

		sqrtNum := math.Sqrt(float64(n))
		for i := 3; i <= int(sqrtNum); i += 2 {
			if n%i == 0 {
				return false
			}
		}

		return true
	}

	num := 2
	for {
		if isPrime(num) {
			primes = append(primes, num)
		}

		if len(primes) == 10001 {
			break
		}

		num++
	}

	fmt.Println(primes[10000])
}
