package array

import (
	"sort"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1 []int
		l2 []int
	}{
		{
			[]int{1, 2, 4},
			[]int{1, 3, 4},
		},
		{
			[]int{1, 2, 4},
			[]int{1, 3, 4},
		},
	}

	for _, test := range tests {
		l1 := createLinklist(test.l1)
		l2 := createLinklist(test.l2)
		head := mergeTwoLists(l1, l2)
		got := linkList2Array(head)

		if !sort.SliceIsSorted(got, func(i, j int) bool {
			return got[i] < got[j]
		}) {
			t.Errorf("test function mergeTwoLists faild, l1:%v,l2:%v, got:%v", test.l1, test.l2, got)
		}
	}
}
