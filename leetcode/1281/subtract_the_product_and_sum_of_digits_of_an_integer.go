package main

func subtractProductAndSum(n int) int {
	p, s := 1, 0
	for n != 0 {
		digit := n % 10
		p *= digit
		s += digit
		n /= 10
	}

	return p - s
}
