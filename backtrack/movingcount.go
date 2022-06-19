package backtrack

type void struct{}

var member void

func movingCount(m int, n int, k int) int {
	sum := func(x int) int {
		s := 0
		for x != 0 {
			s += x % 10
			x = x / 10
		}
		return s
	}
	set := make(map[[2]int]void)

	var dfs func(m, n int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || sum(i)+sum(j) > k {
			return 0
		}
		if _, ok := set[[2]int{i, j}]; ok {
			return 0
		}
		set[[2]int{i, j}] = member
		return 1 + dfs(i+1, j) + dfs(i, j+1)
	}

	return dfs(0, 0)
}
