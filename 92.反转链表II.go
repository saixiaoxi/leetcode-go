package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// 找到 left 之前的节点
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// 开始反转的节点
	start := prev.Next
	then := start.Next

	// 逐个反转节点
	for i := 0; i < right-left; i++ {
		start.Next = then.Next
		then.Next = prev.Next
		prev.Next = then
		then = start.Next
	}

	return dummy.Next
}
