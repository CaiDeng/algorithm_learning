package array

import "testing"

func TestGetInterSectionNode(t *testing.T) {
	tests := []struct {
		l1     []int
		l2     []int
		common []int
	}{
		{
			[]int{4, 1},
			[]int{5, 0, 1},
			[]int{8, 4, 5},
		},
		{
			[]int{0, 9, 1},
			[]int{3},
			[]int{2, 4},
		},
		{
			[]int{2, 6, 4},
			[]int{1, 5},
			[]int{},
		},
	}

	for _, test := range tests {
		l1 := createLinklist(test.l1)
		l2 := createLinklist(test.l2)
		common := createLinklist(test.common)
		cur := l1
		for cur != nil {
			if cur.Next == nil {
				break
			}
			cur = cur.Next
		}
		if cur != nil {
			cur.Next = common
		}

		cur = l2
		for cur != nil {
			if cur.Next == nil {
				break
			}
			cur = cur.Next
		}
		if cur != nil {
			cur.Next = common
		}
		got := getIntersectionNode(l1, l2)
		if got != common {
			t.Errorf("test function getInterSectionNode faild, l1:%v,l2:%v,common:%v", test.l1, test.l2, test.common)
		}

	}
}
