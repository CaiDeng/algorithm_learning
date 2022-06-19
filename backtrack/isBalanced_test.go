package backtrack

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"true",
			args{NewTree([]int{3, 9, 20, invalidVal, invalidVal, 15, 7})},
			true,
		},
		{
			"false",
			args{NewTree([]int{1, 2, 2, 3, 3, invalidVal, invalidVal, 4, 4})},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
