package dynamicplanning

import (
	"algorithm/helper"
	"math"
)

func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return math.MinInt
	}
	rows, columns := len(grid)+1, len(grid[0])+1
	dp := make([][]int, rows)
	row := make([]int, rows*columns)
	for i := range dp {
		dp[i], row = row[:columns], row[columns:]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = grid[i-1][j-1] + helper.Max(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[rows-1][columns-1]
}
