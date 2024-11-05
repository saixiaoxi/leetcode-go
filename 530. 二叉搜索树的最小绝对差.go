package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return -1
	}

	minDiff := math.MaxInt64
	//var prev *TreeNode // prev := &TreeNode{}  这样会直接初始化一个空的实例，指针不为nil
	prev := new(TreeNode)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 中序遍历左子树
		inorder(node.Left)

		// 处理当前节点
		if prev != nil {
			diff := node.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = node

		// 中序遍历右子树
		inorder(node.Right)
	}

	inorder(root)
	return minDiff

}
