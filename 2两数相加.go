package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ls := &ListNode{}
	p := ls
	carry := 0
	for carry > 0 || l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
	}
	return ls.Next // 返回结果链表的头节点
}
