package backtrack

import "testing"

func TestMaxdepth(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			[]int{3, 9, 20, invalidVal, invalidVal, 15, 7},
			3,
		},
	}

	for _, test := range tests {
		root := NewTree(test.input)
		got := maxDepth(root)
		if got != test.want {
			t.Errorf("test function maxDepth failed, intpu:%v, want:%v, got:%v",
				test.input, test.want, got)
		}
	}
}
