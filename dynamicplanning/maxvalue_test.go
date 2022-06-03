package dynamicplanning

import "testing"

func TestMaxValue(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
	}{
		{
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			12,
		},
	}

	for _, test := range tests {
		got := maxValue(test.input)
		if got != test.want {
			t.Errorf("test function maxValue failed, input:%v, want:%v ,real:%v", test.input, test.want, got)
		}
	}
}
