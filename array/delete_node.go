package array

func deleteNode(head *ListNode, val int) *ListNode {
	pre := &ListNode{Next: head}
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			if cur == head {
				return cur.Next
			}
			return head
		}
		pre = cur
		cur = cur.Next
	}
	return head
}
