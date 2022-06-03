package backtrack

import "algorithm/helper"

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

func mirrorTreeRecursion(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	sta := helper.NewStack()
	sta.Push(root)

	for !sta.Empty() {
		top := sta.Pop()
		node := top.(*TreeNode)
		if node.Right != nil {
			sta.Push(node.Right)
		}

		if node.Left != nil {
			sta.Push(node.Left)
		}

		node.Left, node.Right = node.Right, node.Left
	}

	return root
}
