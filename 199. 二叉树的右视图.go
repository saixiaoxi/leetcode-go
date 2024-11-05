package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 记录每一层的最右侧节点值
			if i == levelSize-1 {
				ans = append(ans, node.Val)
			}

			// 将左子节点和右子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return ans
}
