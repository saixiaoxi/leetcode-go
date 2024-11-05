package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowZero, colZero := false, false

	// 检查第一行是否有零
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			rowZero = true
			break
		}
	}

	// 检查第一列是否有零
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true
			break
		}
	}

	// 使用第一行和第一列作为标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据标记将对应的行和列置零
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// 根据标记将对应的列置零
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// 如果第一行需要置零
	if rowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// 如果第一列需要置零
	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
