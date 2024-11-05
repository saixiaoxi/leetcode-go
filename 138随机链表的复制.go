package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 创建一个哈希表来存储原节点和新节点的对应关系
	nodeMap := make(map[*Node]*Node)

	// 第一次遍历，创建新节点，并存储在哈希表中
	current := head
	for current != nil {
		nodeMap[current] = &Node{Val: current.Val}
		current = current.Next
	}

	// 第二次遍历，设置新节点的 next 和 random 指针
	current = head
	for current != nil {
		if current.Next != nil {
			nodeMap[current].Next = nodeMap[current.Next]
		}
		if current.Random != nil {
			nodeMap[current].Random = nodeMap[current.Random]
		}
		current = current.Next
	}

	// 返回新链表的头节点
	return nodeMap[head]
}
