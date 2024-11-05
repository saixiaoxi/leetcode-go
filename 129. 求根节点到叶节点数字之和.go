package main

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	// 更新当前路径的数字
	currentSum = currentSum*10 + node.Val

	// 如果是叶节点，返回当前路径的数字
	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	// 递归计算左子树和右子树的路径和
	leftSum := dfs(node.Left, currentSum)
	rightSum := dfs(node.Right, currentSum)

	// 返回左右子树路径和的总和
	return leftSum + rightSum
}
