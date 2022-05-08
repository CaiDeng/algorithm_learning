package search

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, 0

	for row >= 0 && col < len(matrix[0]) {
		if target > matrix[row][col] {
			col++
		} else if target < matrix[row][col] {
			row--
		} else {
			return true
		}
	}
	return false
}
