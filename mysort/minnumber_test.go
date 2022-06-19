package mysort

import "testing"

func TestMinNumber(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{
			[]int{10, 2},
			"102",
		},
		{
			[]int{3, 30, 34, 5, 9},
			"3033459",
		},
	}

	for _, test := range tests {
		got := minNumber(test.input)
		if got != test.want {
			t.Errorf("test function minNumber failed, input:%v, want:%v,got:%v",
				test.input, test.want, got)
		}
	}
}
