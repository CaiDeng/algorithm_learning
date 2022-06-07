package dynamicplanning

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			"abcabcbb",
			3,
		},
	}

	for _, test := range tests {
		got := lengthOfLongestSubstring(test.input)
		if got != test.want {
			t.Errorf("test function lengthOfLongestSubstring failed, input:%v, want:%v ,real:%v", test.input, test.want, got)
		}
	}
}
