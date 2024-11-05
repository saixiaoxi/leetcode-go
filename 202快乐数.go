package main

// 辅助函数，计算一个数每个位置上的数字的平方和
func getNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// 判断一个数是否为快乐数
func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for n != 1 && !contains(seen, n) {
		seen[n] = struct{}{}
		n = getNext(n)
	}
	return n == 1
}

// 检查集合中是否包含某个元素
func contains(set map[int]struct{}, key int) bool {
	_, exists := set[key]
	return exists
}
