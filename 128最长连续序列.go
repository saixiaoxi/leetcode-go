package main

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	count := 1
	max := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		}
		if nums[i+1] == nums[i]+1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 1
		}
	}
	if count > max {
		max = count
	}

	return max
}
