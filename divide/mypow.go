package divide

import (
	"math"

	"github.com/EngoEngine/math/imath"
)

func myPowReur(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	a := myPowReur(x, n/2)
	a = a * a * math.Pow(x, float64(n&1*n/imath.Abs(n)))
	return a
}

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	res := 1.0
	for n != 0 {
		res *= math.Pow(x, float64(n&1*n/imath.Abs(n)))
		x *= x
		n /= 2
	}

	return res
}
