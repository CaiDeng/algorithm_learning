package search

import "testing"

func TestMinArray(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			[]int{3, 4, 5, 1, 2},
			1,
		},
		{
			[]int{2, 2, 2, 0, 1},
			0,
		},
	}

	for _, test := range tests {
		got := minArray(test.input)
		if test.want != got {
			t.Errorf("test fucntion minArray failed, input:%v, want:%v, real:%v", test.input, test.want, got)
		}
	}
}
