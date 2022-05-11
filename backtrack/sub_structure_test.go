package backtrack

import "testing"

func TestIsSubStructure(t *testing.T) {
	tests := []struct {
		A    []int
		B    []int
		want bool
	}{
		{
			[]int{1, 2, 3},
			[]int{3, 1},
			false,
		},
		{
			[]int{3, 4, 5, 1, 2},
			[]int{4, 1},
			true,
		},
	}

	for _, test := range tests {
		Atree := NewTree(test.A)
		Btree := NewTree(test.B)
		got := isSubStructure(Atree, Btree)
		if got != test.want {
			t.Errorf("test function isSubStructure failed, A:%v, B:%v, want:%v, real:%v", test.A, test.B, test.want, got)
		}
	}
}
