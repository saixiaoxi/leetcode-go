package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		second = second.Next
		first = first.Next
	}

	second.Next = second.Next.Next

	return dummy.Next
}
