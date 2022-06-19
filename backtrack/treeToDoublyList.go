package backtrack

import "algorithm/helper"

func treeToDoublyList(root *TreeNode) (head *TreeNode) {
	head = new(TreeNode)
	p := head
	cur := root
	stack := helper.NewStack()
	for cur != nil {
		stack.Push(cur)
		cur = cur.Left
	}

	var node *TreeNode
	for !stack.Empty() {
		top := stack.Pop()
		node = top.(*TreeNode)
		p.Right = node
		node.Left = p
		p = node
		cur = node.Right

		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
	}
	head.Left = node
	if node != nil {
		node.Right = head.Right
		head.Right.Left = node
	}
	return head
}
