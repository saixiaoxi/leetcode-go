package main

func twoSum(numbers []int, target int) []int {
	start := 1
	end := len(numbers)
	for start < end {
		if numbers[start]+numbers[end] > target {
			end--
		} else if numbers[start]+numbers[end] < target {
			start++
		} else {
			return []int{start, end}
		}
	}
	return []int{-1}
}
