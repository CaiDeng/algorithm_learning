package divide

import (
	"math"
	"testing"
)

const tollerate float64 = 1.0e-16

func Test_myPowReur(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			"x <0,n <0",
			args{
				-4.3,
				-5,
			},
			1.0,
		},
		{
			"x <0, n = 0",
			args{
				3.2,
				0,
			},
			1.1,
		},
		{
			"x <0, n >0",
			args{
				-2.1,
				7,
			},
			1.2,
		},
		{
			"x=0,n >=0",
			args{
				0,
				9,
			},
			2.1,
		},
		{
			"x>0, n < 0",
			args{
				3.1,
				-13,
			},
			2.1,
		}, {
			"x > 0, n = 0",
			args{
				2.4,
				0,
			},
			2.1,
		}, {
			"x >0 ,n >0",
			args{
				4.3,
				11,
			},
			1.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, fuzzGot := myPowReur(tt.args.x, tt.args.n), math.Pow(tt.args.x, float64(tt.args.n)); math.Abs(got-fuzzGot) > tollerate {
				t.Errorf("myPowReur() = %v, want %v", got, fuzzGot)
			}
		})
	}
}

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			"x <0,n <0",
			args{
				-4.3,
				-5,
			},
			1.0,
		},
		{
			"x <0, n = 0",
			args{
				3.2,
				0,
			},
			1.1,
		},
		{
			"x <0, n >0",
			args{
				-2.1,
				7,
			},
			1.2,
		},
		{
			"x=0,n >=0",
			args{
				0,
				9,
			},
			2.1,
		},
		{
			"x>0, n < 0",
			args{
				3.1,
				-13,
			},
			2.1,
		}, {
			"x > 0, n = 0",
			args{
				2.4,
				0,
			},
			2.1,
		}, {
			"x >0 ,n >0",
			args{
				4.3,
				11,
			},
			1.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, fuzzGot := myPow(tt.args.x, tt.args.n), math.Pow(tt.args.x, float64(tt.args.n)); math.Abs(got-fuzzGot) > tollerate {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
