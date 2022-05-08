package search

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	i, j := 0, len(numbers)-1
	for i < j {
		mid := (i + j) / 2
		if numbers[mid] < numbers[j] {
			j = mid
		} else if numbers[mid] > numbers[j] {
			i = mid + 1
		} else {
			j--
		}
	}
	return numbers[i]
}
