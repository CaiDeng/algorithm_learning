package array

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]&1 == 1 {
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
		if nums[j]&1 == 0 {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	return nums
}
