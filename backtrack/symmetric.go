package backtrack

func symmetricTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return symmetricTree(left.Left, right.Right) && symmetricTree(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmetricTree(root.Left, root.Right)
}
