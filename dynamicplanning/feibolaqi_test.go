package dynamicplanning

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{2, 1},
		{5, 5},
	}

	for _, test := range tests {
		got := fib(test.input)
		if got != test.want {
			t.Errorf("test function Fib failed, input:%v, want:%v, real:%v", test.input, test.want, got)
		}
	}
}
