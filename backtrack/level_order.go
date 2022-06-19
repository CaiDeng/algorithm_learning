package backtrack

import "algorithm/helper"

const invalidVal = -1

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Search 查找树节点
func (t *TreeNode) Search(val int) *TreeNode {
	if t == nil {
		return nil
	}

	cur := t
	for cur != nil {
		if cur.Val == val {
			return cur
		}

		if val < t.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return nil

}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	q := helper.NewQueue()
	q.Enqueue(root)

	for !q.Empty() {
		levelLen := q.Size()
		levelList := []int{}
		for i := 0; i < levelLen; i++ {
			peek, _ := q.Dequeue()
			node := peek.(*TreeNode)
			levelList = append(levelList, node.Val)
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		result = append(result, levelList)
	}
	return result
}
