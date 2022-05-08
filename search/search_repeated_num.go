package search

type searchMode int

const (
	left searchMode = iota
	right
)

func search(nums []int, target int) int {
	leftIndex, find := binarySearch(nums, target, left)
	if !find {
		return 0
	}
	rightIndex, find := binarySearch(nums, target, right)
	if !find {
		return 0
	}

	return rightIndex - leftIndex - 1
}

func binarySearch(nums []int, target int, mod searchMode) (index int, find bool) {
	index = -1
	if len(nums) == 0 {
		return
	}

	find = false
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			find = true
			if mod == left {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
	}
	if mod == left {
		index = j
	} else {
		index = i
	}

	return
}
