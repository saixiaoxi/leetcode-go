package main

func containsNearbyDuplicate(nums []int, k int) bool {
	maps := make(map[int]int)
	for index, num := range nums {
		if _, ok := maps[num]; ok {
			if index-maps[num] <= k {
				return true
			} else {
				maps[num] = index
			}
		} else {
			maps[num] = index
		}
	}
	return false
}
