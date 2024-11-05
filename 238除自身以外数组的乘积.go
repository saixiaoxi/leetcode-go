package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 初始化前缀积数组
	prefix := make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	// 初始化后缀积数组
	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	// 计算结果数组
	for i := 0; i < n; i++ {
		answer[i] = prefix[i] * suffix[i]
	}

	return answer
}
