package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 递归翻转左子树和右子树
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	// 交换左子树和右子树
	root.Left = right
	root.Right = left
	return root
}
