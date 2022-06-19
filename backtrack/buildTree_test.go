package backtrack

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preOrder []int
		inOrder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"多节点",
			args{
				[]int{3, 9, 20, 15, 7},
				[]int{9, 3, 15, 20, 7},
			},
			[]int{3, 9, 20, 15, 7},
		},
		{
			"单结点",
			args{
				[]int{-1},
				[]int{-1},
			},
			[]int{-1},
		},
		{
			"没有右子树",
			args{
				[]int{1, 2, 3, 4},
				[]int{3, 2, 4, 1},
			},
			[]int{1, 2, 3, 4},
		},
		{
			"没有左子树",
			args{
				[]int{1, 2, 3, 4},
				[]int{1, 3, 2, 4},
			},
			[]int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preOrder, tt.args.inOrder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
