package main

func canConstruct(ransomNote string, magazine string) bool {
	maps := make(map[rune]int)
	for _, info := range magazine {
		if _, ok := maps[info]; ok {
			maps[info]++
			continue
		} else {
			maps[info] = 1
		}
	}
	for _, info := range ransomNote {
		if maps[info] > 0 {
			maps[info]--
			continue
		} else {
			return false
		}
	}
	return true
}
