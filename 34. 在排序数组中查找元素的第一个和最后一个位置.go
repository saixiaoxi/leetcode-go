package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if nums[left] < target {
				left = mid
			} else if nums[right] > target {
				right = mid
			} else if nums[left] == target && nums[right] == target {
				// 检查边界条件
				if (left == 0 || nums[left-1] < target) && (right == len(nums)-1 || nums[right+1] > target) {
					return []int{left, right}
				}
				if left > 0 && nums[left-1] == target {
					left--
				}
				if right < len(nums)-1 && nums[right+1] == target {
					right++
				}
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}
