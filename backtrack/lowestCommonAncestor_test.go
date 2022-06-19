package backtrack

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"normal",
			args{
				root: NewTree([]int{6, 2, 8, 0, 4, 7, 9, invalidVal, invalidVal, 3, 5}),
				p:    2,
				q:    8,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.root.Search(tt.args.p), tt.args.root.Search(tt.args.q)); !reflect.DeepEqual(got, tt.args.root.Search(tt.want)) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
			got := lowestCommonAncestor(tt.args.root, tt.args.root.Search(tt.args.p), tt.args.root.Search(tt.args.q))
			fuzzeGot := lowestCommonAncestorRecur(tt.args.root, tt.args.root.Search(tt.args.p), tt.args.root.Search(tt.args.q))
			if !reflect.DeepEqual(got, fuzzeGot) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
