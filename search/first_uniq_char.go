package search

func firstUniqChar(s string) byte {

	charMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if charMap[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}
