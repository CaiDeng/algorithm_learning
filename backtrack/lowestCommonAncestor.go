package backtrack

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == nil || q == nil {
		return nil
	}

	cur := root
	for cur != nil {
		if p == cur || q == cur {
			return cur
		}

		if (p.Val-cur.Val)*(q.Val-cur.Val) < 0 {
			return cur
		}

		if p.Val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return cur
}

func lowestCommonAncestorRecur(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p == nil || q == nil {
		return nil
	}

	if p == root || q == root {
		return root
	}

	if (p.Val < root.Val && q.Val > root.Val) || (p.Val > root.Val && q.Val < root.Val) {
		return root
	}

	if p.Val < root.Val {
		return lowestCommonAncestorRecur(root.Left, p, q)
	}
	return lowestCommonAncestorRecur(root.Right, p, q)

}
