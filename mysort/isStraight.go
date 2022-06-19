package mysort

import "math"

type void struct{}

var member void

func isStraight(nums []int) bool {
	length := len(nums)
	set := make(map[int]void)
	max := math.MinInt
	min := math.MaxInt
	for _, v := range nums {
		if v == 0 {
			continue
		}
		if _, ok := set[v]; ok {
			return false
		}
		set[v] = member
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	if max-min < length && max-min >= 0 {
		return true
	}
	return false
}
