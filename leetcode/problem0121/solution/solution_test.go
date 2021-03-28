package solution

import "testing"

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		output := maxProfit(test.input)
		t.Logf("input: %v, output %v, expected: %v", test.input, output, test.output)
		if output != test.output {
			t.Errorf("Failed")
		}
	}
}
