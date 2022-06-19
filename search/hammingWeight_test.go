package search

import (
	"math"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"num=0",
			args{0},
			0,
		},
		{
			"num=13",
			args{13},
			3,
		},
		{
			"num=17",
			args{17},
			2,
		},
		{
			"num=max",
			args{math.MaxUint32},
			32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
