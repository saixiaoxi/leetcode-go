package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归展开左子树和右子树
	flatten(root.Left)
	flatten(root.Right)

	// 保存右子树
	rightSubtree := root.Right

	// 将左子树插入到右子树的位置
	root.Right = root.Left
	root.Left = nil

	// 找到当前右子树的末端
	current := root
	for current.Right != nil {
		current = current.Right
	}

	// 将保存的右子树连接到当前右子树的末端
	current.Right = rightSubtree
}
