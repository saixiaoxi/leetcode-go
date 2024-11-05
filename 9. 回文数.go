package main

func isPalindrome(x int) bool {

	// 负数和以0结尾但不为0的数不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除中间的数字
	return x == revertedNumber || x == revertedNumber/10
}
