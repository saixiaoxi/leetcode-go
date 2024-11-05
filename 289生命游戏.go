package main

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	// 创建一个副本
	copyBoard := make([][]int, m)
	for i := range board {
		copyBoard[i] = make([]int, n)
		copy(copyBoard[i], board[i])
	}

	// 定义方向数组，用于遍历八个相邻位置
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// 遍历每个细胞
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := 0
			// 计算周围八个位置的活细胞数
			for _, dir := range directions {
				r, c := i+dir[0], j+dir[1]
				if r >= 0 && r < m && c >= 0 && c < n && copyBoard[r][c] == 1 {
					liveNeighbors++
				}
			}

			// 根据规则更新细胞状态
			if copyBoard[i][j] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[i][j] = 0
			}
			if copyBoard[i][j] == 0 && liveNeighbors == 3 {
				board[i][j] = 1
			}
		}
	}
}
