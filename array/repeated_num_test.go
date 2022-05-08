package array

import (
	"algorithm/helper"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	tests := []struct {
		params []int
		want   []int
	}{
		{[]int{2, 3, 1, 0, 2, 5, 3}, []int{2, 3}},
	}

	for _, test := range tests {
		got := findRepeatNumber(test.params)
		if helper.ContainsInt(test.want, got) == -1 {
			t.Errorf("tese functions FindRepeatNumber failed, params:%v, wants:%v, real:%v", test.params,
				test.want, got)
		}
	}
}
