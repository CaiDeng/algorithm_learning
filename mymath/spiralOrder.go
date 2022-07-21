package mymath

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, len(matrix)*len(matrix[0]))
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1

	index := 0
	for {
		//  从左往右
		for i := l; i <= r; i++ {
			res[index] = matrix[t][i]
			index++
		}

		if t++; t > b {
			break
		}

		// 从上往下
		for i := t; i <= b; i++ {
			res[index] = matrix[i][r]
			index++
		}

		if r--; l > r {
			break
		}

		// 从右往左
		for i := r; i >= l; i-- {
			res[index] = matrix[b][i]
			index++
		}
		if b--; t > b {
			break
		}

		for i := b; i >= t; i-- {
			res[index] = matrix[i][l]
			index++
		}
		if l++; l > r {
			break
		}
	}
	return res
}
