package helper

// ContainsInt slice contain the value
func ContainsInt(array []int, val int) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// Min two integer number minmum
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max two integer number maxmum
func Max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
