package main

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的
// 子序列
// 。
// 输入：nums = [10,9,2,5,3,4,101,18]
// 1 1 1 2 2 3 4 4
// 输出：4
// 解释：最长递增子序列是 [2,3,4,101]，因此长度为 4 。

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLength := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
