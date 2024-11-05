package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// 将字符串转换为小写
	s = strings.ToLower(s)

	// 移除所有非字母数字字符
	var filteredString strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			filteredString.WriteRune(char)
		}
	}

	// 获取处理后的字符串
	filtered := filteredString.String()

	// 检查是否为回文串
	length := len(filtered)
	for i := 0; i < length/2; i++ {
		if filtered[i] != filtered[length-1-i] {
			return false
		}
	}
	return true
}
