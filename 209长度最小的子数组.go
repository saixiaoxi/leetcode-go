package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLength := n
	start := 0
	end := 0
	sum := 0
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			if end-start+1 < minLength {
				minLength = end - start + 1
			}
			sum -= nums[start]
			start++
		}

		end++
	}
	start--
	if sum+nums[start] < target {
		return 0
	}

	return minLength
}
