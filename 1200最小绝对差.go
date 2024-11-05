package main

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var minarr [][]int
	if len(arr) < 2 {
		return minarr // 如果数组长度小于2，直接返回空切片
	}
	min := arr[1] - arr[0] // 初始化min为第一对元素的差值
	for i := 0; i < len(arr)-1; i++ {
		mid := arr[i+1] - arr[i]
		if mid < min {
			min = mid
			minarr = [][]int{{arr[i], arr[i+1]}} // 重新初始化minarr并添加当前元素对
		} else if mid == min {
			minarr = append(minarr, []int{arr[i], arr[i+1]}) // 添加当前元素对
		}
	}
	return minarr
}
