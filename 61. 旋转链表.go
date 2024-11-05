package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 计算链表的长度
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length += 1
	}

	// 将链表连接成环
	tail.Next = head

	// 计算新的头节点的位置
	k = k % length
	stepsToNewHead := length - k

	// 找到新的头节点和尾节点
	newTail := tail
	for i := 0; i < stepsToNewHead; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next

	// 断开环
	newTail.Next = nil

	return newHead
}
