package mystrings

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"the sky is blue",
			"blue is sky the",
		},
		{
			"  hello world!  ",
			"world! hello",
		},
	}

	for _, test := range tests {
		got := reverseWords(test.input)
		if got != test.want {
			t.Errorf("test function reversWords failed, input:%v,want:%v,got:%v", test.input, test.want, got)
		}
	}
}
