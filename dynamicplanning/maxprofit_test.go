package dynamicplanning

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, test := range tests {
		got := maxProfit(test.input)
		if got != test.want {
			t.Errorf("test function maxProfit failed, input:%v, want:%v, real:%v", test.input, test.want, got)
		}
	}
}
