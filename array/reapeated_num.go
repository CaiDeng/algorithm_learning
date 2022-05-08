package array

func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 0; i < len(nums); {
		if i == nums[i] {
			i += 1
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		} else {
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}
