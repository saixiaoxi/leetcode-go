package main

func strStr(haystack string, needle string) int {
	for start := 0; start < len(haystack); start++ {
		if haystack[start] == needle[0] {
			end := 0
			for i := 0; i < len(needle) && start+i < len(haystack); i++ {
				end = start + i
				if haystack[start+i] != needle[i] {
					break
				}
			}
			var str string
			if str = haystack[start : end+1]; str == needle {
				return start
			}
		}
	}
	return -1
}
