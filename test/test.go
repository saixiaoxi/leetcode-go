package main

import "fmt"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	// nums1 := []int{3, 4, 5, 1, 2}
	// fmt.Println(findMin(nums1))
	// nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	// fmt.Println(findMin(nums2))

	channel := make(chan int)

	channel <- 1
	close(channel)
	fmt.Println("wuhu") // 输出: 1

}
