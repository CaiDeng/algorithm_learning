package backtrack

func recur(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	} else if a == nil {
		return false
	} else if a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}

	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
