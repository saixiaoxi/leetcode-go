package main

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return 1
	}
	candys := make([]int, len(ratings))
	for i := range candys {
		candys[i] = 1
	}
	if ratings[0] > ratings[1] {
		candys[0]++
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] && candys[i] <= candys[i-1] {
			candys[i] = candys[i-1] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candys[i-1] <= candys[i] {
			candys[i-1] = candys[i] + 1
		}
	}

	candy := 0
	for i := range candys {
		candy += candys[i]
	}
	return candy
}
