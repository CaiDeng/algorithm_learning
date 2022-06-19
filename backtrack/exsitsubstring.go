package backtrack

const isFound byte = ' '

func exist(board [][]byte, word string) bool {
	depth := len(word) - 1
	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if k > depth {
			return true
		}
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == isFound {
			return false
		}

		found := false
		dx, dy := []int{0, 0, -1, 1}, []int{-1, 1, 0, 0}
		if board[i][j] == word[k] {
			tmp := board[i][j]
			board[i][j] = isFound
			for n := 0; n < len(dx); n++ {
				found = dfs(i+dx[n], j+dy[n], k+1)
				if found {
					break
				}
			}
			board[i][j] = tmp
		}
		return found
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
