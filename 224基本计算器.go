package main

import "unicode"

func calculate(s string) int {
	stack := []int{}
	result := 0
	number := 0
	sign := 1

	for i := 0; i < len(s); i++ {
		char := s[i]

		if unicode.IsDigit(rune(char)) {
			number = 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				number = number*10 + int(s[i]-'0')
				i++
			}
			i-- // 回退一位，因为外层循环还会增加一次
			result += sign * number
		} else if char == '+' {
			sign = 1
		} else if char == '-' {
			sign = -1
		} else if char == '(' {
			stack = append(stack, result)
			stack = append(stack, sign)
			result = 0
			sign = 1
		} else if char == ')' {
			result *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return result
}
