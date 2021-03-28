package solution

import "testing"

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
	for _, test := range tests {
		output := rob(test.input)
		t.Logf("input: %v, output %v, expected: %v", test.input, output, test.output)
		if output != test.output {
			t.Errorf("Failed")
		}
	}
}
