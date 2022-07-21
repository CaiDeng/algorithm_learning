package mymath

func constructArr(a []int) []int {
	dp := 1
	res := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		res[i] = dp
		dp *= a[i]
	}

	dp = 1
	for i := len(a) - 1; i >= 0; i-- {
		res[i] *= dp
		dp *= a[i]
	}
	return res
}
