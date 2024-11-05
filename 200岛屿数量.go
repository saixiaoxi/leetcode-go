package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var dfs func(int, int)
	dfs = func(i, j int) {
		// 越界或者当前位置不是陆地时返回
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		// 将当前位置标记为已访问
		grid[i][j] = '0'
		// 访问当前位置的上下左右四个位置
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				// 每发现一个岛屿，岛屿数量加一
				count++
				// 然后通过dfs将这个岛屿的所有部分都标记为已访问
				dfs(i, j)
			}
		}
	}

	return count
}
