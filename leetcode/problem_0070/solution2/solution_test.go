package solution

import "testing"

func TestProblem(t *testing.T) {
	tests := []struct {
		input, output int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	for _, test := range tests {
		output := climbStairs(test.input)
		t.Logf("input: %v, output %v", test.input, output)
		if output != test.output {
			t.Errorf("expected: %v, output: %v", test.output, output)
		}
	}
}
