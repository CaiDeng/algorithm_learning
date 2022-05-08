package array

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		operations []string
		params     []int
		wants      []int
	}{
		{
			[]string{"construct", "push", "push", "push", "min", "pop", "top", "min"},
			[]int{-1, -2, 0, -3, -1, -1, -1, -1},
			[]int{-1, -1, -1, -1, -3, -1, 0, -2},
		},
	}

	for _, test := range tests {
		var testMinStack MinStack
		got := []int{}
		for index, operation := range test.operations {
			switch operation {
			case "construct":
				testMinStack = Constructor()
				got = append(got, -1)
			case "push":
				testMinStack.Push(test.params[index])
				got = append(got, -1)
			case "min":
				got = append(got, testMinStack.Min())
			case "pop":
				testMinStack.Pop()
				got = append(got, -1)
			case "top":
				got = append(got, testMinStack.Top())

			}
		}

		if !Equal(got, test.wants) {
			t.Errorf("MinStack test fail: operations:%v, params:%v, wants:%v, real:%v", test.operations, test.params, test.wants, got)
		}
	}

}
