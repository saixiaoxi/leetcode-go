package main

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	res := nums[0]
	min := 0
	sum := 0
	sum_all := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		sum_all += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
		if sum > res {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	sum = 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum < min {
			min = sum
		}
		if sum > 0 {
			sum = 0
		}
	}
	if sum_all != min && sum_all-min > res {
		return sum_all - min
	}

	return res
}
