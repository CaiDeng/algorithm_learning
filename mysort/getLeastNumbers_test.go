package mysort

import (
	"testing"

	"github.com/thoas/go-funk"
)

func TestGetLeastNumbers(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  []int
	}{
		{
			[]int{3, 2, 1},
			2,
			[]int{1, 2},
		},
		{
			[]int{0, 1, 2, 1},
			1,
			[]int{0},
		},
		{
			[]int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4},
			10,
			[]int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4},
		},
	}

	for _, test := range tests {
		got := getLeastNumbers(test.input, test.k)
		if !funk.Equal(got, test.want) {
			t.Errorf("test function getLeastNumber failed, input:%v, k:%v,want:%v,got:%v",
				test.input, test.k, test.want, got)
		}
	}
}
