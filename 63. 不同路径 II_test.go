package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Seconded struct {
	obstacleGrid [][]int
	expected     int
}

func TestUniquePathsWithObstacles(t *testing.T) {
	unique := make([]Seconded, 4) // 初始化切片长度为 4

	unique[0] = struct {
		obstacleGrid [][]int
		expected     int
	}{
		obstacleGrid: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		expected: 2,
	}

	unique[1] = struct {
		obstacleGrid [][]int
		expected     int
	}{
		obstacleGrid: [][]int{
			{0, 1},
			{0, 0},
		},
		expected: 1,
	}

	unique[2] = struct {
		obstacleGrid [][]int
		expected     int
	}{
		obstacleGrid: [][]int{
			{1, 0},
			{0, 0},
		},
		expected: 0,
	}

	unique[3] = struct {
		obstacleGrid [][]int
		expected     int
	}{
		obstacleGrid: [][]int{
			{0, 0},
			{0, 1},
		},
		expected: 0,
	}

	for _, test := range unique {
		result := uniquePathsWithObstacles(test.obstacleGrid)
		assert.Equal(t, test.expected, result)
	}
}
