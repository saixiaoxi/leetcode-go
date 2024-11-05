package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 && obstacleGrid[0][j-1] == 1 {
			obstacleGrid[0][j] = 1
		} else {
			obstacleGrid[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				if obstacleGrid[i-1][j] != 0 {
					obstacleGrid[i][j] += obstacleGrid[i-1][j]
				}
				if obstacleGrid[i][j-1] != 0 {
					obstacleGrid[i][j] += obstacleGrid[i][j-1]
				}
			} else {
				obstacleGrid[i][j] = 0
			}

		}
	}
	return obstacleGrid[m-1][n-1]
}
