package backtrack

import "algorithm/helper"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
