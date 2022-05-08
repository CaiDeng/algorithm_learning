package array

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	old2CopyMap := map[*Node]*Node{}
	for p := head; p != nil; p = p.Next {
		old2CopyMap[p] = new(Node)
		old2CopyMap[p].Val = p.Val
	}
	old2CopyMap[nil] = nil

	copyHead := old2CopyMap[head]

	for old_p, p := head, copyHead; old_p != nil; old_p, p = old_p.Next, p.Next {
		p.Next = old2CopyMap[old_p.Next]
		p.Random = old2CopyMap[old_p.Random]
	}

	return copyHead
}

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 复制节点，1->1'->2->2'->nil
	for p := head; p != nil; p = p.Next.Next {
		copyNode := new(Node)
		*copyNode = *p
		p.Next = copyNode
	}
	// 处理复制节点的随机指针
	copyHead := head.Next
	p := copyHead
	for ; p != nil && p.Next != nil; p = p.Next.Next {
		randomPointer := p.Random
		if randomPointer != nil {
			p.Random = randomPointer.Next
		}
	}
	randomPointer := p.Random
	if randomPointer != nil {
		p.Random = randomPointer.Next
	}

	// 分离复制链表和原始链表
	p = head
	for ; p != nil && p.Next.Next != nil; p = p.Next {
		p.Next.Next, p.Next = p.Next.Next.Next, p.Next.Next
	}
	p.Next = nil

	return copyHead

}
