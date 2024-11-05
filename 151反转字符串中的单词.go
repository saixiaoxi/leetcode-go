package main

import (
	"strings"
)

func reverseWords(s string) string {
	newS := make([]string, 0)
	start := len(s) - 1
	end := len(s) - 1
	for start >= 0 {
		// 跳过空格
		for start >= 0 && s[start] == ' ' {
			start--
			end--
		}
		// 找到一个单词的结尾
		for start >= 0 && s[start] != ' ' {
			start--
		}
		// 将单词添加到 newS 中
		if start < end {
			newS = append(newS, s[start+1:end+1])
		}
		// 更新 end
		end = start
	}

	str := strings.Join(newS, " ")
	return str
}
