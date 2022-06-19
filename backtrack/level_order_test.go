package backtrack

import (
	"algorithm/helper"
	"testing"

	funk "github.com/thoas/go-funk"
)

// NewTree 给定一个完全树的列表，构建成一个树结构
func NewTree(nodeList []int) (root *TreeNode) {
	if len(nodeList) == 0 {
		return nil
	}

	q := helper.NewQueue()
	index := 0
	root = &TreeNode{Val: nodeList[index]}
	q.Enqueue(root)
	parents := []*TreeNode{}
	for !q.Empty() {
		length := q.Size()
		nextParent := []*TreeNode{}
		for i := 0; i < length; i++ {
			peek, _ := q.Dequeue()
			node := peek.(*TreeNode)
			if len(parents) > 0 {
				parent := parents[i>>1]
				if i&1 == 0 {
					parent.Left = node
				} else {
					parent.Right = node
				}
			}

			if node != nil {
				nextParent = append(nextParent, node)
				var subNode *TreeNode
				index++
				if index < len(nodeList) {
					if nodeList[index] != invalidVal {
						subNode = &TreeNode{Val: nodeList[index]}
					}
					q.Enqueue(subNode)
				}
				index++
				if index < len(nodeList) {
					subNode = nil
					if nodeList[index] != invalidVal {
						subNode = &TreeNode{Val: nodeList[index]}
					}
					q.Enqueue(subNode)
				}
			}
		}

		parents = nextParent
	}

	return root
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
