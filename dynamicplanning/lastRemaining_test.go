package dynamicplanning

import "testing"

func Test_lastRemaining(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"n=1,m=1",
			args{1, 1},
			0,
		},
		{
			"n=5,m=3",
			args{5, 3},
			3,
		},
		{
			"n=10,m=17",
			args{10, 17},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastRemaining(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("lastRemaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
