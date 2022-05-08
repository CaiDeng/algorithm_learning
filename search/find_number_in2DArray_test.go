package search

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			[][]int{
				[]int{1, 4, 7, 11, 15},
				[]int{2, 5, 8, 12, 19},
				[]int{3, 6, 9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			},
			5,
			true,
		},
		{
			[][]int{
				[]int{1, 4, 7, 11, 15},
				[]int{2, 5, 8, 12, 19},
				[]int{3, 6, 9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			},
			20,
			false,
		},
	}

	for _, test := range tests {
		got := findNumberIn2DArray(test.matrix, test.target)
		if got != test.want {
			t.Errorf("test function findNumberIn2DArray failed, matrix:%v, target:%v, want:%v, real:%v",
				test.matrix, test.target, test.want, got)
		}
	}
}
