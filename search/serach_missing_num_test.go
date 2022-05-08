package search

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0, 1, 3}, 2},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 9}, 8},
		{[]int{1}, 0},
		{[]int{0}, 1},
	}

	for _, test := range tests {
		got := missingNumber(test.input)
		if got != test.want {
			t.Errorf("test function missingNumber failed, input:%v, want:%v, real:%v", test.input, test.want, got)
		}
	}
}
