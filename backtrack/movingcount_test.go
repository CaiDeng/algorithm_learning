package backtrack

import "testing"

func TestMovingCount(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		k    int
		want int
	}{
		{
			2, 3, 1, 3,
		},
		{
			3, 1, 0, 1,
		},
	}

	for _, test := range tests {
		got := movingCount(test.m, test.n, test.k)
		if got != test.want {
			t.Errorf("test function movingcount failed, m:%v, n:%v, k:%v, got:%v", test.m, test.n, test.k, got)
		}
	}
}
