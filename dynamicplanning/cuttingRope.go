package dynamicplanning

import (
	"math"

	"github.com/EngoEngine/math/imath"
)

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		max := math.MinInt
		for j := 1; j <= i/2; j++ {
			tmp := imath.Max(dp[j], j) * imath.Max(dp[i-j], i-j)
			if tmp > max {
				max = tmp
			}
		}
		dp[i] = max
	}
	return dp[n]
}
