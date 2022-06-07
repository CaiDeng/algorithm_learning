package array

// Node Definition for singly-linked list
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinklist(list []int) *ListNode {
	preHead := &ListNode{}
	pre := preHead
	for _, val := range list {
		cur := new(ListNode)
		cur.Val = val
		pre.Next = cur
		pre = cur
	}
	pre.Next = nil
	return preHead.Next
}

func linkList2Array(head *ListNode) []int {
	cur := head
	list := []int{}
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}
	return list
}
