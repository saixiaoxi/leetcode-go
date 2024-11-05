package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	//将排序后的每个字符串作为键，这样就能保证异位词的键相等，不需要遍历处理。
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

// sortString 对字符串中的字符进行排序
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
