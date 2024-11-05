package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 首先对数组进行排序
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复的数字
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum < 0:
				left++
			case sum > 0:
				right--
			default:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// 跳过重复的数字
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}

// func threeSum(nums []int) [][]int {
//     sort.Ints(nums) // 首先对数组进行排序
//     var result [][]int

//     for i := 0; i < len(nums)-2; i++ {
//         // 跳过重复的数字
//         if i > 0 && nums[i] == nums[i-1] {
//             continue
//         }

//         target := -nums[i]
//         numMap := make(map[int]bool)
//         for j := i + 1; j < len(nums); j++ {
//             if _, found := numMap[target-nums[j]]; found {
//                 result = append(result, []int{nums[i], nums[j], target - nums[j]})
//                 // 跳过重复的数字
//                 for j+1 < len(nums) && nums[j] == nums[j+1] {
//                     j++
//                 }
//             }
//             numMap[nums[j]] = true
//         }
//     }

//     return result
// }
