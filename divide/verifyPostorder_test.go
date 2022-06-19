package divide

import "testing"

func Test_verifyPostorder(t *testing.T) {
	type args struct {
		postOrder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"false",
			args{
				[]int{1, 6, 3, 2, 5},
			},
			false,
		},
		{
			"true",
			args{
				[]int{1, 3, 2, 6, 5},
			},
			true,
		},
		{
			"no right subtree but ture",
			args{
				[]int{1, 3, 2, 6, 7},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postOrder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
