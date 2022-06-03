package dynamicplanning

import (
	"algorithm/helper"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	cost := prices[0]
	for i := 1; i < len(prices); i++ {
		cost = helper.Min(cost, prices[i])
		profit = helper.Max(profit, prices[i]-cost)
	}

	return profit
}
