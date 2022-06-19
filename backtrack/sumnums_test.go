package backtrack

import "testing"

func Test_sumNums(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"n>1,n=10",
			args{10},
			55,
		},
		{
			"n<=0",
			args{0},
			0,
		},
		{
			"n=1",
			args{1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNums(tt.args.n); got != tt.want {
				t.Errorf("sumNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
