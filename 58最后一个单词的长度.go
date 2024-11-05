package main

func lengthOfLastWord(s string) int {
	left, right := len(s)-1, len(s)-1
	for {
		if s[left] == ' ' {
			left--
			right--
		} else {
			break
		}
	}
	for s[left] != ' ' && left > 0 {
		if left > 1 && s[left-1] == ' ' {
			break
		}
		left--
	}
	return right - left + 1
}
