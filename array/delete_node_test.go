package array

import (
	"testing"

	funk "github.com/thoas/go-funk"
)

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		input []int
		val   int
		want  []int
	}{
		{
			[]int{4, 5, 1, 9},
			5,
			[]int{4, 1, 9},
		},
		{
			[]int{4, 5, 1, 9},
			1,
			[]int{4, 5, 9},
		},
		{
			[]int{4, 5, 1, 9},
			6,
			[]int{4, 5, 1, 9},
		},
	}

	for _, test := range tests {
		linkList := createLinklist(test.input)
		head := deleteNode(linkList, test.val)
		got := linkList2Array(head)
		if !funk.Equal(got, test.want) {
			t.Errorf("test function deleteNode failed, input:%v,val:%v, want:%v,got:%v",
				test.input, test.val, test.want, got)
		}
	}

}
