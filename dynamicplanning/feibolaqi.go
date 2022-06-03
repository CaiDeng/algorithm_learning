package dynamicplanning

func fib(n int) int {
	a0, a1 := 0, 1
	for i := 0; i < n; i++ {
		a0, a1 = a1, (a0+a1)%1000000007
	}

	return a0

}
