package bit

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"a<0,b <0",
			args{-3, -4},
			-7,
		},
		{
			"a<0,b=0",
			args{-5, 0},
			-5,
		},
		{
			"a < 0 ,b > 0",
			args{-5, 5},
			0,
		},
		{
			"a=0,b<0",
			args{0, -3},
			-3,
		},
		{
			"a=0,b=0",
			args{0, 0},
			0,
		},
		{
			"a=0,b>0",
			args{0, 7},
			7,
		},
		{
			"a>0, b <0",
			args{4, -5},
			-1,
		},
		{
			"a >0, b = 0",
			args{8, 0},
			8,
		},
		{
			"a > 0, b >0",
			args{5, 6},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
