package backtrack

import "github.com/EngoEngine/math/imath"

func isBalanced(root *TreeNode) bool {
	var depth func(*TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := depth(root.Left)
		if leftDepth == -1 {
			return -1
		}
		rightDepth := depth(root.Right)
		if rightDepth == -1 {
			return -1
		}

		if imath.Abs(leftDepth-rightDepth) > 1 {
			return -1
		}
		return imath.Max(leftDepth, rightDepth) + 1
	}

	if depth(root) == -1 {
		return false
	}
	return true
}
