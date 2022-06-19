package backtrack

import "algorithm/helper"

func kthLargest(root *TreeNode, k int) int {
	stack := helper.NewStack()
	cur := root
	for cur != nil {
		stack.Push(cur)
		cur = cur.Right
	}

	res := invalidVal
	for i := 0; i < k && !stack.Empty(); i++ {
		top := stack.Pop()
		node := top.(*TreeNode)
		res = node.Val
		cur = node.Left
		for cur != nil {
			stack.Push(cur)
			cur = cur.Right
		}
	}

	return res
}
