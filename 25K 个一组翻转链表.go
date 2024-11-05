package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	// 这里在头部自己加一个节点
	dummy := &ListNode{0, head}
	pre := dummy

	for {
		end := pre
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start
	}
}

func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
