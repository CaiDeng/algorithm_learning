package array

import (
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestCqueue(t *testing.T) {
	tests := []struct {
		operations []string
		params     []int
		wants      []int
	}{
		{
			[]string{"CQueue", "appendTail", "deleteHead", "deleteHead"},
			[]int{-1, 3, -1, -1},
			[]int{-1, -1, 3, -1},
		},
		{
			[]string{"CQueue", "deleteHead", "appendTail", "appendTail", "deleteHead", "deleteHead"},
			[]int{-1, -1, 5, 2, -1, -1},
			[]int{-1, -1, -1, -1, 5, 2},
		},
	}

	for _, test := range tests {
		got := make([]int, 0)
		var testCqueue CQueue
		for index, operation := range test.operations {
			switch operation {
			case "CQueue":
				testCqueue = NewCQueue()
				got = append(got, -1)
			case "appendTail":
				testCqueue.AppendTail(test.params[index])
				got = append(got, -1)
			case "deleteHead":
				got = append(got, testCqueue.DeleteHead())
			}
		}

		if !Equal(got, test.wants) {
			t.Errorf("Cqueue test failed, operations:%v, prams:%v, wants:%v, real: %v", test.operations,
				test.params, test.wants, got)
		}
	}
}
