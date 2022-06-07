package array

import (
	"testing"

	"github.com/thoas/go-funk"
)

func TestGetKthFromEnd(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			7,
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, 5},
			0,
			[]int{},
		},
	}

	for _, test := range tests {
		linkList := createLinklist(test.input)
		head := getKthFromEnd(linkList, test.k)
		got := linkList2Array(head)

		if !funk.Equal(got, test.want) {
			t.Errorf("test function failed, input:%v,k:%v, want:%v, real:%v", test.input, test.k, test.want, got)
		}
	}
}
