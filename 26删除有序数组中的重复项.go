package main

func removeDuplicates(nums []int) int {
	newIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[newIndex] = nums[i]
			newIndex++
		}
	}
	return newIndex
}
