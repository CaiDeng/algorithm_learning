package dynamicplanning

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	charIndexMap := map[byte]int{}
	dp := 0
	maxLength := 0

	for j := 0; j < len(s); j++ {
		i, ok := charIndexMap[s[j]]
		if !ok {
			i = -1
		}
		charIndexMap[s[j]] = j
		if dp < j-i {
			dp++
		} else {
			dp = j - i
		}

		if dp > maxLength {
			maxLength = dp
		}
	}
	return maxLength
}
