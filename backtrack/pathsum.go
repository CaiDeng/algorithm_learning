package backtrack

func pathSum(root *TreeNode, target int) [][]int {
	var dfs func(root *TreeNode, target int)
	res := [][]int{}
	path := []int{}
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Val == target && root.Left == nil && root.Right == nil {
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			res = append(res, copyPath)
		}
		target -= root.Val
		dfs(root.Left, target)
		dfs(root.Right, target)
		path = path[:len(path)-1]
		return
	}

	dfs(root, target)
	return res
}
