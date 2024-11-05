package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// 初始化第一列
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = min(grid[i-1][j]+grid[i][j], grid[i][j-1]+grid[i][j])
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
