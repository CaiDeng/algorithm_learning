package array

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return nil
	}
	former, latter := head, head
	i := 0
	for ; i < k && former != nil; i++ {
		former = former.Next
	}
	if i < k {
		return nil
	}

	for former != nil {
		latter = latter.Next
		former = former.Next
	}
	return latter
}
