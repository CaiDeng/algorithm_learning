package strings

import "testing"

func TestReverseLeftWords(t *testing.T) {
	tests := []struct {
		params string
		pos    int
		wants  string
	}{
		{"abcdefg", 2, "cdefgab"},
		{"lrloseumgh", 6, "umghlrlose"},
	}

	for _, test := range tests {
		got := reverseLeftWords(test.params, test.pos)
		if got != test.wants {
			t.Errorf("tese functions reverseLeftWords failed, params:%v, pos:%v, wants:%v, real:%v", test.params,
				test.pos, test.wants, got)
		}
	}
}
