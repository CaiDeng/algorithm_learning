package search

import (
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		input string
		want  byte
	}{
		{"abaccdeff", 'b'},
		{"", ' '},
	}

	for _, test := range tests {
		got := firstUniqChar(test.input)
		if got != test.want {
			t.Errorf("test fucntion minArray failed, input:%v, want:%v, real:%v", test.input, test.want, got)
		}
	}
}
