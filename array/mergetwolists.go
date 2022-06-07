package array

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := new(ListNode)
	pre := preHead
	curL1 := l1
	curL2 := l2
	for curL1 != nil && curL2 != nil {
		if curL1.Val < curL2.Val {
			pre.Next = curL1
			curL1 = curL1.Next
		} else {
			pre.Next = curL2
			curL2 = curL2.Next
		}
		pre = pre.Next
	}

	if curL1 == nil {
		pre.Next = curL2
	}

	if curL2 == nil {
		pre.Next = curL1
	}

	return preHead.Next
}
