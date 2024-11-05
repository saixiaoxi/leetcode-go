package main

// 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获 所有 被围绕的区域：

// 连接：一个单元格与水平或垂直方向上相邻的单元格连接。
// 区域：连接所有 'O' 的单元格来形成一个区域。
// 围绕：如果您可以用 'X' 单元格 连接这个区域，并且区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕。
// 通过将输入矩阵 board 中的所有 'O' 替换为 'X' 来 捕获被围绕的区域。

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m, n := len(board), len(board[0])

	// 定义 DFS 函数
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = '#' // 标记为特殊字符
		dfs(i-1, j)       // 上
		dfs(i+1, j)       // 下
		dfs(i, j-1)       // 左
		dfs(i, j+1)       // 右
	}

	// 遍历边界，标记不被围绕的 'O'
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}

	// 遍历整个矩阵，进行替换
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X' // 被围绕的 'O' 替换为 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O' // 还原之前标记的 'O'
			}
		}
	}
}
