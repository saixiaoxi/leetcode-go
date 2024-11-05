package main

func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1 // 初始化为1，表示和为0的子数组默认有一个（空数组）

	for _, num := range nums {
		sum += num // 累加当前元素到总和中
		if _, ok := prefixSum[sum-k]; ok {
			count += prefixSum[sum-k] // 如果存在一个之前的前缀和，使得当前前缀和减去它等于k，则找到了一个符合条件的子数组
		}
		prefixSum[sum]++ // 更新前缀和出现的次数
	}

	return count
}
