package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var s []string
	start := nums[0]

	for i := 0; i < len(nums); i++ {
		// 如果是最后一个元素，或者当前元素与下一个元素不连续
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if start == nums[i] {
				s = append(s, strconv.Itoa(start))
			} else {
				s = append(s, fmt.Sprintf("%d->%d", start, nums[i]))
			}
			// 更新 start 为下一个元素
			if i != len(nums)-1 {
				start = nums[i+1]
			}
		}
	}

	return s
}
