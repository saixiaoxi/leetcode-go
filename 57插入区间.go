package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0

	// 添加所有在 newInterval 之前且不重叠的区间
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 合并所有与 newInterval 重叠的区间
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	// 添加剩余的区间
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}

// 获取两个整数的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 获取两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
