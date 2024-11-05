package main

import (
	"fmt"
)

func maxPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	if len(points) == 1 {
		return 1
	}

	maxCount := 1

	for i := 0; i < len(points); i++ {
		slopes := make(map[string]int)
		duplicate := 0
		curMax := 0

		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 && dy == 0 {
				duplicate++
				continue
			}

			gcd := getGCD(dx, dy)
			dx /= gcd
			dy /= gcd

			slope := fmt.Sprintf("%d/%d", dx, dy)
			slopes[slope]++
			if slopes[slope] > curMax {
				curMax = slopes[slope]
			}
		}

		maxCount = max(maxCount, curMax+duplicate+1)
	}

	return maxCount
}

func getGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return getGCD(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
