package main

import "strings"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// 初始化结果数组，长度为两个字符串长度之和，最大可能的长度
	result := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			// 计算当前位的乘积，并加上之前这一位的结果，因为可能有进位
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			// 加上之前这一位的结果
			sum := mul + result[i+j+1]
			// 更新当前位的结果
			result[i+j+1] = sum % 10
			// 处理进位
			result[i+j] += sum / 10
		}
	}
	// 转换结果数组为字符串
	var sb strings.Builder
	// 跳过结果数组前面的0
	i := 0
	for i < len(result) && result[i] == 0 {
		i++
	}
	for ; i < len(result); i++ {
		sb.WriteByte(byte(result[i] + '0'))
	}
	return sb.String()
}
