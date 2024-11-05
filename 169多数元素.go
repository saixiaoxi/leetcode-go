package main

import "sort"

func majorityElement(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	sort.Ints(nums)
	max := 1
	mid := 1
	maxNum := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			mid++
			if i == len(nums)-1 {
				if max < mid {
					max = mid
					mid = 1
					maxNum = nums[i-1]
				}
			}
		} else {
			if max < mid {
				max = mid
				mid = 1
				maxNum = nums[i-1]
			}
		}

	}
	return maxNum
}
