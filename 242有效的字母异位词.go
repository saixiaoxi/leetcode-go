package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	maps := make(map[byte]int)
	mapt := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		maps[s[i]]++
		mapt[t[i]]++
	}

	for k, v := range maps {
		if mapt[k] != v {
			return false
		}
	}

	return true
}
