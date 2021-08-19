package main

func maxProfit(prices []int) int {
	profit := 0

	for idx := 0; idx < len(prices)-1; idx++ { 		// до n-1, потому что в последний день нельзя покупать ничего
		if prices[idx+1] > prices[idx] { 			// если завтра акции дороже, то купить их
			profit += prices[idx+1] - prices[idx]
		}
	}

	return profit
}
