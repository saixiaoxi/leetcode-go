package main

// // ListNode 定义链表节点
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// sortList 对链表进行排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针找到链表的中间节点
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 将链表分成两半
	prev.Next = nil

	// 递归排序两半
	l1 := sortList(head)
	l2 := sortList(slow)

	// 合并排序后的两半
	return merge(l1, l2)
}

// merge 合并两个有序链表
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}
