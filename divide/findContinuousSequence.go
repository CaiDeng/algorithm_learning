package divide

func findContinuousSequence(target int) [][]int {
	res := [][]int{}
	newSequence := func(i, j int) []int {
		r := make([]int, j-i+1)
		for index := range r {
			r[index] = i
			i++
		}
		return r
	}
	for i, j := 1, 2; i <= target/2 && i < j; {
		sequenceSum := (i + j) * (j - i + 1) / 2
		if sequenceSum == target {
			res = append(res, newSequence(i, j))
			i++
		} else if sequenceSum > target {
			i++
		} else {
			j++
		}
	}
	return res
}
