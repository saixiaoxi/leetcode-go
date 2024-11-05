package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		// 获取结果数组中最后一个区间
		last := result[len(result)-1]
		current := intervals[i]

		// 如果当前区间的起始位置小于等于最后一个区间的结束位置，说明有重叠
		if current[0] <= last[1] {
			// 更新最后一个区间的结束位置
			last[1] = max(last[1], current[1])
			result[len(result)-1] = last
		} else {
			// 否则，将当前区间加入结果数组
			result = append(result, current)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
