package array

import "testing"

func convertLinkList(input [][]int) (head *Node) {
	head = nil
	if len(input) == 0 {
		return
	}
	// 创建节点，并建立（节点索引，节点指针）的映射表
	index2NodeMap := []*Node{}
	pre_p := &Node{
		0,
		nil,
		nil,
	}
	head = pre_p
	for _, node := range input {
		tmp := new(Node)
		tmp.Val = node[0]
		pre_p.Next = tmp
		pre_p = pre_p.Next
		index2NodeMap = append(index2NodeMap, tmp)
	}
	pre_p.Next = nil

	// 处理随机指针
	p := head.Next
	for _, node := range input {
		randomNodeIndex := node[1]
		if randomNodeIndex == -1 {
			p.Random = nil
		} else {
			p.Random = index2NodeMap[randomNodeIndex]
		}
		p = p.Next
	}

	head = head.Next
	return
}

func convertDoubleSlice(head *Node) (result [][]int) {
	result = nil
	if head == nil {
		return result
	}

	index2NodeMap := map[*Node]int{nil: -1}
	for index, p := 0, head; p != nil; p, index = p.Next, index+1 {
		result = append(result, []int{p.Val, -1})
		index2NodeMap[p] = index
	}

	// 处理随机指针
	for index, p := 0, head; p != nil && index < len(result); p, index = p.Next, index+1 {
		result[index][1] = index2NodeMap[p.Random]
	}
	return result
}

func EqualDoubleSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for index, value := range a {
		if !Equal(value, b[index]) {
			return false
		}
	}
	return true
}

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		input [][]int
		wants [][]int
	}{
		{
			[][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
			[][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		{
			[][]int{{1, 1}, {2, 1}},
			[][]int{{1, 1}, {2, 1}},
		},
	}

	for _, test := range tests {
		head := convertLinkList(test.input)
		gotHead := copyRandomList1(head)
		got := convertDoubleSlice(gotHead)
		if !EqualDoubleSlice(test.wants, got) {
			t.Errorf("test copyRandomList failed. input:%v, wants:%v, real:%v", test.input, test.wants, got)
		}
	}
}
