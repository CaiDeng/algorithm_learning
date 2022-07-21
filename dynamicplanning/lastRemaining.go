package dynamicplanning

func lastRemaining(n int, m int) int {
	dp := 0

	for i := 2; i <= n; i++ {
		dp = (dp + m) % i
	}
	return dp

}
