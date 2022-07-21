package bit

func singleNumbers(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	m := 1
	for xor&m == 0 {
		m <<= 1
	}

	x, y := 0, 0
	for _, v := range nums {
		if v&m != 0 {
			x ^= v
		} else {
			y ^= v
		}
	}
	return []int{x, y}
}
