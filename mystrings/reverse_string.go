package mystrings

import "bytes"

func reverseLeftWords(s string, n int) string {
	var buffer bytes.Buffer
	if n <= 0 || n > len(s) {
		return ""
	}
	sRune := []rune(s)

	for i := n; i < len(sRune)+n; i++ {
		buffer.WriteString(string(sRune[i%len(sRune)]))
	}

	return buffer.String()

}
