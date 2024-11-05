package main

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	newIndex := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[newIndex-2] {
			nums[newIndex] = nums[i]
			newIndex++
		}
	}

	return newIndex
}
