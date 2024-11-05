package main

func firstMissingPositive(nums []int) int {
	maps := make(map[int]int)
	for index, num := range nums {
		maps[num] = index
	}
	for i := 1; ; i++ {
		if _, ok := maps[i]; ok {
			continue
		}
		return i
	}
}

// func firstMissingPositive(nums []int) int {
//     n := len(nums)

//     // 将每个正整数 num 放到索引 num-1 的位置上
//     for i := 0; i < n; i++ {
//         for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
//             nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
//         }
//     }

//     // 找到第一个不满足 nums[i] == i + 1 的位置
//     for i := 0; i < n; i++ {
//         if nums[i] != i + 1 {
//             return i + 1
//         }
//     }

//     // 如果所有位置都满足 nums[i] == i + 1，则返回 n + 1
//     return n + 1
// }
