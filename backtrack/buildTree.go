package backtrack

func buildTree(preOrder, inOrder []int) []int {

	root2Index := map[int]int{}
	for index, v := range inOrder {
		root2Index[v] = index
	}

	var build func(int, int, int, int) *TreeNode
	// 递归参数是当前树的前序和中序的首尾索引，主要是为了能够快速找到中序中头节点的索引
	build = func(preOrderHead, preOrderTail, inOrderHead, inOrderTail int) *TreeNode {
		if preOrderTail-preOrderHead < 0 || inOrderTail-inOrderHead < 0 {
			return nil
		}
		if preOrderTail-preOrderHead != inOrderTail-inOrderHead {
			panic("params error,the number of preorder must be equal to inorder")
		}

		root := &TreeNode{Val: preOrder[preOrderHead]}
		if preOrderTail-preOrderHead == 0 {
			return root
		}

		rootInOrder := root2Index[preOrder[preOrderHead]]
		leftPreOrderHead := preOrderHead + 1
		leftPreOrderTail := preOrderHead + rootInOrder - inOrderHead
		leftInOrderTail := rootInOrder - 1
		leftInorDerHead := inOrderHead
		root.Left = build(leftPreOrderHead, leftPreOrderTail, leftInorDerHead, leftInOrderTail)

		rightPreOrderHead := leftPreOrderTail + 1
		rightPreOrderTail := preOrderTail
		rightInOrderHead := rootInOrder + 1
		rightInOrderTail := inOrderTail
		root.Right = build(rightPreOrderHead, rightPreOrderTail, rightInOrderHead, rightInOrderTail)
		return root
	}

	root := build(0, len(preOrder)-1, 0, len(inOrder)-1)

	var inOrderIter func(*TreeNode)
	result := []int{}
	inOrderIter = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		result = append(result, tn.Val)
		inOrderIter(tn.Left)
		inOrderIter(tn.Right)
	}

	inOrderIter(root)
	return result
}
