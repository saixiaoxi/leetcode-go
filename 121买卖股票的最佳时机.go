package main

func maxProfit(prices []int) int {
	min := prices[0]
	mid := 0
	max := 0
	for _, price := range prices {
		if price < min {
			min = price
		}
		if price > min {
			mid = price - min
			if mid > max {
				max = mid
			}
		}
	}
	return max
}
