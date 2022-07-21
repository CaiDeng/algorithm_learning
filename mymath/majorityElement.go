package mymath

func majorityElement(nums []int) int {
	vote, major := 0, 0
	for _, v := range nums {
		if vote == 0 {
			major = v
		}
		if v == major {
			vote++
		} else {
			vote--
		}
	}
	return major
}
