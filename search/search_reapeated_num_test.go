package search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, 2},
		{[]int{5, 7, 7, 8, 8, 10}, 6, 0},
	}

	for _, test := range tests {
		got := search(test.nums, test.target)
		if got != test.want {
			t.Errorf("test function search failed, nums:%v,target:%v,want:%v, real:%v", test.nums, test.target, test.want, got)
		}

	}
}
