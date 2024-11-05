package main

import (
	"container/heap"
)

// Pair 定义数对
type Pair struct {
	num1, num2 int
}

// MinHeap 定义最小堆
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].num1+h[i].num2 < h[j].num1+h[j].num2 }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	h := &MinHeap{}
	heap.Init(h)

	// 初始化最小堆
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, Pair{nums1[i], nums2[0]})
	}

	result := [][]int{}
	for k > 0 && h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		result = append(result, []int{p.num1, p.num2})
		k--

		// 找到下一个数对
		if p.num2 == nums2[len(nums2)-1] {
			continue
		}
		heap.Push(h, Pair{p.num1, nums2[findIndex(nums2, p.num2)+1]})
	}

	return result
}

// findIndex 查找元素在数组中的索引
func findIndex(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	result := [][]int{}
	m, n := len(nums1), len(nums2)
	indices := make([]int, m)

	for k > 0 {
		minSum := int(^uint(0) >> 1) // 初始化为最大整数
		x := -1

		for i := 0; i < m; i++ {
			if indices[i] < n && nums1[i]+nums2[indices[i]] < minSum {
				minSum = nums1[i] + nums2[indices[i]]
				x = i
			}
		}

		if x == -1 {
			break
		}

		result = append(result, []int{nums1[x], nums2[indices[x]]})
		indices[x]++
		k--
	}

	return result
}
