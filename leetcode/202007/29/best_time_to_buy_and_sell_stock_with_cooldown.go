package main

import "math"

func maxProfit(prices []int) int {
	sold := int(math.MinInt64)
	held := int(math.MinInt64)
	reset := 0

	for _, price := range prices {
		preSold := sold
		sold = held + price
		held = max(held, reset-price)
		reset = max(reset, preSold)
	}

	return max(sold, reset)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
