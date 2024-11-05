package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	// 将整数数组转换为字符串数组
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	// 自定义排序规则
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})

	// 如果最大的数字是0，直接返回"0"
	if strNums[0] == "0" {
		return "0"
	}

	// 合并字符串数组为一个字符串
	return strings.Join(strNums, "")
}
