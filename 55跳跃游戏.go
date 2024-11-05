package main

func canJump(nums []int) bool {
	maxReach := 0
	for i, num := range nums {
		// 如果当前位置超过了之前能到达的最远距离，则无法到达最后
		if i > maxReach {
			return false
		}
		// 更新能到达的最远距离
		if i+num > maxReach {
			maxReach = i + num
		}
		// 如果最远距离已经超过或达到最后一个下标，则可以到达
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}
