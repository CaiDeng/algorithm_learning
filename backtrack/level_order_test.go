package backtrack

import (
	"algorithm/helper"
	"testing"

	funk "github.com/thoas/go-funk"
)

const invalidVal = -1

// NewTree 给定一个完全树的列表，构建成一个树结构
func NewTree(nodeList []int) (root *TreeNode) {
	if len(nodeList) == 0 {
		return nil
	}

	pre := new(TreeNode)
	q := helper.NewQueue()
	q.Enqueue(pre)
	for index, val := range nodeList {
		peek, _ := q.Peek()
		parent := peek.(*TreeNode)

		var leafNode *TreeNode
		if val != invalidVal {
			leafNode = &TreeNode{Val: val}
		}

		q.Enqueue(leafNode)

		if parent != nil {
			if index%2 == 1 {
				parent.Left = leafNode
			} else {
				parent.Right = leafNode
			}
		}

		if index%2 == 0 {
			q.Dequeue()
		}

	}
	return pre.Right
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		input []int
		want  [][]int
	}{
		{
			[]int{3, 9, 20, invalidVal, invalidVal, 15, 7},
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}

	for _, test := range tests {
		root := NewTree(test.input)
		got := levelOrder(root)
		if !funk.Equal(got, test.want) {
			t.Errorf("test fucntion levelOrder failed, input:%v, want:%v, real:%v", test.input, test.want, got)
		}
	}
}
