package main

// 到达第 n 阶的方法数等于到达第 n-1 阶和第 n-2 阶的方法数之和。
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		temp := a + b
		a = b
		b = temp
	}
	return b
}
