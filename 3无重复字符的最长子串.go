package main

func lengthOfLongestSubstring(s string) int {
	lastOccurrence := make(map[byte]int)
	maxLength := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if index, ok := lastOccurrence[s[right]]; ok && index >= left {
			left = index + 1
		}
		lastOccurrence[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 动态规划实现
// func lengthOfLongestSubstring(s string) int {
// 	n := len(s)
// 	if n == 0 {
// 		return 0
// 	}

// 	// dp[i] 表示以 s[i] 结尾的最长无重复子串的长度
// 	dp := make([]int, n)
// 	// lastOccurrence 记录字符最后出现的位置
// 	lastOccurrence := make(map[byte]int)

// 	dp[0] = 1
// 	lastOccurrence[s[0]] = 0
// 	maxLength := 1

// 	for i := 1; i < n; i++ {
// 		if index, ok := lastOccurrence[s[i]]; ok {
// 			// 如果字符 s[i] 出现过，更新 dp[i]
// 			dp[i] = min(dp[i-1]+1, i-index)
// 		} else {
// 			// 如果字符 s[i] 没有出现过，dp[i] = dp[i-1] + 1
// 			dp[i] = dp[i-1] + 1
// 		}
// 		// 更新字符 s[i] 的最后出现位置
// 		lastOccurrence[s[i]] = i
// 		// 更新最大长度
// 		maxLength = max(maxLength, dp[i])
// 	}

// 	return maxLength
// }
