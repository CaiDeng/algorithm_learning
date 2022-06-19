package backtrack

import (
	"testing"

	"github.com/thoas/go-funk"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   [][]int
	}{
		{
			[]int{5, 4, 8, 11, invalidVal, 13, 4, 7, 2, invalidVal, invalidVal, 5, 1},
			22,
			[][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
		{
			[]int{1, 2},
			0,
			[][]int{},
		},
		{
			[]int{1, 2, 3},
			5,
			[][]int{},
		},
	}

	for _, test := range tests {
		root := NewTree(test.input)
		got := pathSum(root, test.target)
		if !funk.Equal(got, test.want) {
			t.Errorf("test function pathSum failed, input:%v,target:%v,want:%v,got:%v",
				test.input, test.target, test.want, got)
		}
	}
}
