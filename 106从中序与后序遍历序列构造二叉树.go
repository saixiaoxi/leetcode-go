package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	// 后序遍历的最后一个元素是根节点
	rootVal := postorder[len(postorder)-1]
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
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

	return root
}
