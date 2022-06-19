package backtrack

func sumNums(n int) int {

	res := 0
	var sum func(int) int
	sum = func(n int) int {
		x := n > 1 && sum(n-1) > 0
		res += n
		x = !x
		return res
	}
	sum(n)
	return res
}
