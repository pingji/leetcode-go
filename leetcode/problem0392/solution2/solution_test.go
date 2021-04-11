package solution

import "testing"

func TestProblem(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		output bool
	}{
		{"", "ahbgdc", true},
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"aaaaaa", "bbaaaa", false},
	}
	for index, test := range tests {
		output := isSubsequence(test.s, test.t)
		t.Logf("index: %v, s: %v, t: %v, output %v", index, test.s, test.t, output)
		if output != test.output {
			t.Errorf("Failed！index: %v, expected: %v, output: %v", index, test.output, output)
		}
	}
}
