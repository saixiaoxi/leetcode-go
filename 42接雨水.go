package main

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	rain := 0

	for left < right {
		if maxLeft < maxRight {
			left++
			maxLeft = max(maxLeft, height[left])
			rain += maxLeft - height[left]
		} else {
			right--
			maxRight = max(maxRight, height[right])
			rain += maxRight - height[right]
		}
	}
	return rain
}
