package main

func maxProfit2(prices []int) int {
	profit := 0
	min := prices[0]
	for index, price := range prices {
		if price < min {
			min = price
		}
		if price > min {
			profit += price - min
			min = prices[index]
		}
	}
	return profit
}
