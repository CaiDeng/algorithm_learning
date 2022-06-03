package backtrack

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{
			[]int{1, 2, 2, 3, 4, 4, 3},
			true,
		},
		{
			[]int{1, 2, 2, invalidVal, 3, invalidVal, 3},
			false,
		},
	}

	for _, test := range tests {
		root := NewTree(test.input)
		got := isSymmetric(root)

		if got != test.want {
			t.Errorf("test function isSymmetric failed, input:%v, want:%v, real:%v", test.input, test.want, got)
		}
	}
}
