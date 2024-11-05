package main

func minPathSum(grid [][]int) int {

	minsum := make([][]int, len(grid))
	for i := range minsum {
		minsum[i] = make([]int, len(grid[0]))
	}
	minsum[0][0] = grid[0][0]
	minsum[0][1] = grid[0][0] + grid[0][1]
	minsum[1][0] = grid[0][0] + grid[1][0]
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			minsum[i][j] = min(minsum[i-1][j]+minsum[i][j], minsum[i][j-1]+minsum[i][j])
		}
	}

	return minsum[len(grid)-1][len(grid[0])-1]
}
