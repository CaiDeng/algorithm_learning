package backtrack

import (
	"testing"

	"github.com/thoas/go-funk"
)

func TestTreeToDoublyList(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			[]int{4, 2, 5, 1, 3},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{12, 5, 18, 2, 9, 15, 19, invalidVal, invalidVal, invalidVal, invalidVal, 13, 17},
			[]int{2, 5, 9, 12, 13, 15, 17, 18, 19},
		},
	}

	for _, test := range tests {
		root := NewTree(test.input)
		linkList := treeToDoublyList(root)
		cur := linkList.Right
		got := []int{}
		for cur != nil {
			got = append(got, cur.Val)
			if cur == linkList.Left {
				break
			}
			cur = cur.Right
		}

		if !funk.Equal(got, test.want) {
			t.Errorf("test function treeToDoulbeList failed, input:%v, want:%v,got:%v",
				test.input, test.want, got)
		}
	}
}
