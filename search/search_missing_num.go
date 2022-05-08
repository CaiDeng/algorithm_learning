package search

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return i
}
