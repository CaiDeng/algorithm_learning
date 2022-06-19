package mysort

import "math/rand"

func getLeastNumbers(arr []int, k int) []int {
	sorted := []int{}
	if len(arr) == 0 || k <= 0 {
		return sorted
	}

	partion := func(p, q int) int {
		randIndex := rand.Intn(q-p+1) + p
		arr[q], arr[randIndex] = arr[randIndex], arr[q]
		i, j := p, q-1
		for i <= j {
			if arr[i] < arr[q] {
				i++
			} else if arr[j] >= arr[q] {
				j--
			} else {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[q], arr[i] = arr[i], arr[q]
		return i
	}

	var findk func(int, int) []int
	findk = func(i1, i2 int) []int {
		index := partion(i1, i2)
		if k == index+1 {
			return arr[:k]
		} else if index+1 < k {
			return findk(index+1, i2)
		} else {
			return findk(i1, index-1)
		}
	}
	return findk(0, len(arr)-1)
}
