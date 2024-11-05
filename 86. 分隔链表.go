package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}
	next := head
	len := 1
	for next.Next != nil {
		next = next.Next
		len++
	}
	for i := 0; i < len; i++ {
		if head.Val > x {
			dummy.Next = head.Next
			head.Next = nil
			next.Next = head
			next = next.Next
			head = dummy.Next
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
