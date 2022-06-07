package array

import (
	"testing"

	"github.com/thoas/go-funk"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{2, 7},
		},
		{
			[]int{10, 26, 30, 31, 47, 60},
			40,
			[]int{10, 30},
		},
	}

	for _, test := range tests {
		got := twoSum(test.input, test.target)
		if !funk.Equal(got, test.want) {
			t.Errorf("test function twoSum failed, input:%v,target:%v,want:%v,got:%v", test.input, test.target, test.want, got)
		}
	}
}
