package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// 定义数字到字母的映射
	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var combinations []string
	var combination []byte

	// 回溯函数
	var backtrack func(index int)
	backtrack = func(index int) {
		// 如果当前索引等于输入字符串的长度，表示已经生成了一个完整的组合
		if index == len(digits) {
			combinations = append(combinations, string(combination))
			return
		}

		// 获取当前数字对应的字母
		letters := phoneMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			// 将当前字母添加到组合中
			combination = append(combination, letters[i])
			// 递归生成下一个位置的组合
			backtrack(index + 1)
			// 回溯，移除最后一个添加的字母
			combination = combination[:len(combination)-1]
		}
	}
	backtrack(0)
	return combinations
}
