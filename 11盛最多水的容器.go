package main

// func maxArea(height []int) int {
// 	max := 0
// 	for i := 0; i < len(height); i++ {
// 		for j := i; j < len(height); j++ {
// 			if height[i] < height[j] {
// 				if height[i]*(j-i) > max {
// 					max = height[i] * (j - i)
// 				}
// 			} else {
// 				if height[j]*(j-i) > max {
// 					max = height[j] * (j - i)
// 				}
// 			}
// 		}
// 	}
// 	return max
// }

func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1

	for left < right {
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		area := h * (right - left + 1)
		if area > max {
			max = area
		}
	}

	return max
}
