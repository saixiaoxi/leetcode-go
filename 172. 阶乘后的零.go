package main

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		n /= 5
		count += count
	}
	return count
}
