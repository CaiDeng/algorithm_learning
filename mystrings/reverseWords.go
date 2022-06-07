package mystrings

import "strings"

func reverseWords(s string) string {
	b := []string{}
	s = strings.Trim(s, " ")
	for i, j := len(s)-1, len(s)-1; i >= 0; {
		for ; i >= 0 && s[i] != ' '; i-- {
		}
		b = append(b, s[i+1:j+1])

		for ; i >= 0 && s[i] == ' '; i-- {
		}
		j = i
	}
	return strings.Join(b, " ")
}
