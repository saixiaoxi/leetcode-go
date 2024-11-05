package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	maps := make([][]byte, numRows)
	for i := range maps {
		maps[i] = make([]byte, 0)
	}

	i, step := 0, 1
	for index := 0; index < len(s); index++ {
		maps[i] = append(maps[i], s[index])
		if i == 0 {
			step = 1
		} else if i == numRows-1 {
			step = -1
		}
		i += step
	}

	// 将二维切片转换为字符串
	var result strings.Builder
	for _, row := range maps {
		result.Write(row)
	}

	return result.String()
}
