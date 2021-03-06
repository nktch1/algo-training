package main

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var (
		minPrice  = math.MaxInt32
		maxProfit = 0
	)

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
