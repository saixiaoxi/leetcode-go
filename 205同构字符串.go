package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 创建两个映射，分别用于 s 到 t 和 t 到 s 的字符映射
	mapST := make(map[byte]byte)
	mapTS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		// 检查 s 到 t 的映射
		if mappedChar, ok := mapST[charS]; ok {
			if mappedChar != charT {
				return false
			}
		} else {
			mapST[charS] = charT
		}

		// 检查 t 到 s 的映射
		if mappedChar, ok := mapTS[charT]; ok {
			if mappedChar != charS {
				return false
			}
		} else {
			mapTS[charT] = charS
		}
	}

	return true
}
