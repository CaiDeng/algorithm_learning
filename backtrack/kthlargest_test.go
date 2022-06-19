package backtrack

import "testing"

func TestKthlargest(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  int
	}{
		{
			[]int{3, 1, 4, invalidVal, 2},
			1,
			4,
		},
		{
			[]int{5, 3, 6, 2, 4, invalidVal, invalidVal, 1},
			3,
			4,
		},
	}

	for _, test := range tests {
		root := NewTree(test.input)
		got := kthLargest(root, test.k)
		if got != test.want {
			t.Errorf("test function kthLargest failed, input:%v, k:%v, want:%v,got:%v",
				test.input, test.k, test.want, got)
		}
	}
}
