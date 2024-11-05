package main

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 按照气球的结束位置进行排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		// 如果当前气球的起始位置大于箭的射出位置，需要射出一支新的箭
		if points[i][0] > end {
			arrows++
			end = points[i][1]
		}
	}

	return arrows
}
