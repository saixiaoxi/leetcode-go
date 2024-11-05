package main

import (
	"math"
)

func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	// 记录 t 中每个字符的出现次数
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	// 滑动窗口中的字符计数
	windowCount := make(map[byte]int)
	have, need := 0, len(tCount)
	res := ""
	resLen := math.MaxInt32
	left := 0

	// 滑动窗口右边界
	for right := 0; right < len(s); right++ {
		c := s[right]
		windowCount[c]++

		// 如果当前字符在 t 中，并且窗口中的该字符数量达到了 t 中的数量
		if tCount[c] > 0 && windowCount[c] == tCount[c] {
			have++
		}

		// 当窗口包含 t 中所有字符时，尝试收缩左边界
		for have == need {
			// 更新结果
			if right-left+1 < resLen {
				res = s[left : right+1]
				resLen = right - left + 1
			}

			// 移动左边界
			windowCount[s[left]]--
			if tCount[s[left]] > 0 && windowCount[s[left]] < tCount[s[left]] {
				have--
			}
			left++
		}
	}

	return res
}
