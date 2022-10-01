package main

func maxProfit(prices []int) int {
	toBuy, toSell := 0, 1
	maxProfit := 0

	for toSell < len(prices) {
		if prices[toBuy] < prices[toSell] {
			profit := prices[toSell] - prices[toBuy]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			toBuy = toSell
		}

		toSell++
	}

	return maxProfit
}
