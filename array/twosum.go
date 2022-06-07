package array

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{nums[i], nums[j]}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
	var a byte = ' '
}
