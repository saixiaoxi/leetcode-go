package main

import "math"

// 贪心
func maxSubArray(nums []int) int {
	min := math.MinInt

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > min {
			min = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}
	return min
}

// 动态规划
// func maxSubArray(nums []int) int {
// 	res := math.MinInt

// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		dp[i] = max(nums[i], dp[i-1]+nums[i])
// 		if dp[i] > res {
// 			res = dp[i]
// 		}
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
// 	fmt.Println(maxSubArray(nums))
// }
