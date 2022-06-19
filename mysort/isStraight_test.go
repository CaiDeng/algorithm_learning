package mysort

import "testing"

func TestIsStraight(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{0, 0, 4, 3, 5},
			true,
		},
		{
			[]int{0, 0, 2, 2, 5},
			false,
		},
	}

	for _, test := range tests {
		got := isStraight(test.input)
		if got != test.want {
			t.Errorf("test function isStraight failed, input:%v, want:%v, got:%v",
				test.input, test.want, got)
		}
	}
}
