package main

func twoSum(nums []int, target int) []int {
	// 创建一个哈希表用于存储每个元素及其索引
	numMap := make(map[int]int)
	for i, num := range nums {
		// 检查 target - num 是否已经在哈希表中
		if j, ok := numMap[target-num]; ok {
			// 如果找到了，返回当前元素的索引和哈希表中存储的索引
			return []int{j, i}
		}
		// 如果没有找到，将当前元素及其索引存入哈希表
		numMap[num] = i
	}
	// 如果没有找到符合条件的两个数，返回空切片
	return nil
}
