package dynamicplanning

import (
	"algorithm/helper"
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt
	}

	maxSum := nums[0]
	dpi := nums[0]
	for i := 1; i < len(nums); i++ {
		dpi = helper.Max(dpi+nums[i], nums[i])
		if dpi > maxSum {
			maxSum = dpi
		}
	}
	return maxSum
}
