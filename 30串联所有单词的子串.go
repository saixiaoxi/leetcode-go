package main

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount
	wordMap := make(map[string]int)
	result := []int{}

	// 记录 words 中每个单词的出现次数
	for _, word := range words {
		wordMap[word]++
	}

	// 遍历 s 中每一个可能的起始位置
	for i := 0; i <= len(s)-totalLen; i++ {
		seen := make(map[string]int)
		j := 0
		for ; j < wordCount; j++ {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			if count, ok := wordMap[word]; ok {
				seen[word]++
				if seen[word] > count {
					break
				}
			} else {
				break
			}
		}
		if j == wordCount {
			result = append(result, i)
		}
	}

	return result
}
