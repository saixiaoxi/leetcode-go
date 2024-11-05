package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 先序遍历的第一个元素是根节点
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的位置
	var rootIndex int
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// 递归构造左子树和右子树
	root.Left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root
}
