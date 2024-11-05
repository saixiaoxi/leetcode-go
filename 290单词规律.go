package main

import "strings"

// 用分割函数来分割string即可
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	mapPatternToWord := make(map[byte]string)
	mapWordToPattern := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

		if mappedWord, ok := mapPatternToWord[char]; ok {
			if mappedWord != word {
				return false
			}
		} else {
			mapPatternToWord[char] = word
		}

		if mappedChar, ok := mapWordToPattern[word]; ok {
			if mappedChar != char {
				return false
			}
		} else {
			mapWordToPattern[word] = char
		}
	}

	return true
}
