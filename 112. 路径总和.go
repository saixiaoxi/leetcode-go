package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// 如果当前节点是叶子节点，检查是否满足条件
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// 递归检查左子树和右子树
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
