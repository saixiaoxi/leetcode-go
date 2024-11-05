package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	left, right := 0, 0
	max := 0
	maxstart := 0
	maxend := 0
	for i := 0; i < len(s); i++ {
		left = i
		right = i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			right++
		}
		right--
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		left++
		right--
		if max < right-left {
			max = right - left
			maxstart = left
			maxend = right
		}
	}
	return s[maxstart : maxend+1]
}

// func longestPalindrome(s string) string {
// 	n := len(s)
// 	if n < 2 {
// 		return s
// 	}

// 	// 初始化 dp 数组
// 	dp := make([][]bool, n)
// 	for i := range dp {
// 		dp[i] = make([]bool, n)
// 	}

// 	// 所有长度为 1 的子串都是回文
// 	for i := 0; i < n; i++ {
// 		dp[i][i] = true
// 	}

// 	start, maxLength := 0, 1

// 	// 填充 dp 数组
// 	for length := 2; length <= n; length++ {
// 		for i := 0; i <= n-length; i++ {
// 			j := i + length - 1
// 			if s[i] == s[j] {
// 				if length == 2 {
// 					dp[i][j] = true
// 				} else {
// 					dp[i][j] = dp[i+1][j-1]
// 				}
// 				if dp[i][j] && length > maxLength {
// 					start = i
// 					maxLength = length
// 				}
// 			}
// 		}
// 	}

// 	return s[start : start+maxLength]
// }
