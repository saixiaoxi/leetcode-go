package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// BSTIterator 结构定义
type BSTIterator struct {
	stack []*TreeNode
}

// Constructor 初始化 BSTIterator 类
func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	iterator.pushAllLeft(root)
	return iterator
}

// Next 将指针向右移动，然后返回指针处的数字
func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.pushAllLeft(node.Right)
	return node.Val
}

// HasNext 如果向指针右侧遍历存在数字，则返回 true；否则返回 false
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// pushAllLeft 将所有左子节点压入栈中
func (this *BSTIterator) pushAllLeft(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}
