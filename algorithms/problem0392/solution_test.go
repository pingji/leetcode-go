package leetcode

import "testing"

func TestProblem(t *testing.T) {
	tests := []struct {
		s  string
		t string
		output bool
	}{
		{"", "ahbgdc", true},
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, test := range tests {
		output := isSubsequence(test.s, test.t)
		t.Logf("s: %v, t: %v, output %v, expected: %v", test.s, test.t, output, test.output)
		if output != test.output {
			t.Errorf("Failed")
		}
	}
}