package backtrack

import "testing"

func TestMirrorTree(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			[]int{4, 7, 2, 9, 6, 3, 1},
		},
	}

	for _, test := range tests {
		root := NewTree(test.input)
		root = mirrorTree(root)
		want := NewTree(test.want)

		if !isSubStructure(root, want) || !isSubStructure(want, root) {
			t.Errorf("test function MirrorTree failed, input:%v, want:%v", test.input, test.want)
		}

	}
}

func TestMirrorTreeRecursion(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			[]int{4, 2, 7, 1, 3, 6, 9},
			[]int{4, 7, 2, 9, 6, 3, 1},
		},
	}

	for _, test := range tests {
		root := NewTree(test.input)
		root = mirrorTreeRecursion(root)
		want := NewTree(test.want)

		if !isSubStructure(root, want) || !isSubStructure(want, root) {
			t.Errorf("test function MirrorTree failed, input:%v, want:%v", test.input, test.want)
		}

	}
}
