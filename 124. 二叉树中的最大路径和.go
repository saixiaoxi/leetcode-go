package main

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子树的最大贡献值
		leftGain := max(0, maxGain(node.Left))
		rightGain := max(0, maxGain(node.Right))

		// 当前节点的最大路径和
		currentPathSum := node.Val + leftGain + rightGain

		// 更新全局最大路径和
		maxSum = max(maxSum, currentPathSum)

		// 返回当前节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}
