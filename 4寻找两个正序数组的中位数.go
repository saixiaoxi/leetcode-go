package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mid float64
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 0 {
		mid = float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	} else {
		mid = float64(nums1[(len(nums1)-1)/2])
	}
	return mid
}
