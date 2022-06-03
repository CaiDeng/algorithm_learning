package dynamicplanning

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
	}

	for _, test := range tests {
		got := maxSubArray(test.input)
		if got != test.want {
			t.Errorf("test function maxSubArray failed, input:%v, want:%v, real:%v", test.input, test.want, got)
		}
	}
}
