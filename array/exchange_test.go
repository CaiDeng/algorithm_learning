package array

import (
	"testing"

	"github.com/thoas/go-funk"
)

func TestExchange(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 3, 2, 4},
		},
	}

	for _, test := range tests {
		got := exchange(test.input)
		if !funk.Equal(got, test.want) {
			t.Errorf("test function exchange failed,input:%v,want:%v,got:%v", test.input, test.want, got)
		}
	}
}
