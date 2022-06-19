package backtrack

import "container/list"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() != 0 {
		depth++
		levelLength := queue.Len()
		for i := 0; i < levelLength; i++ {
			front := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			if front.Left != nil {
				queue.PushBack(front.Left)
			}

			if front.Right != nil {
				queue.PushBack(front.Right)
			}
		}
	}
	return depth
}
