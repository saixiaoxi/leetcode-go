package main

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		// 如果当前节点和下一个节点的值相同
		if head.Next != nil && head.Val == head.Next.Val {
			// 跳过所有相同的节点
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			// 将 prev.Next 指向不同的节点
			prev.Next = head.Next
		} else {
			// 如果当前节点和下一个节点的值不同，移动 prev
			prev = prev.Next
		}
		// 移动 head
		head = head.Next
	}

	return dummy.Next
}
