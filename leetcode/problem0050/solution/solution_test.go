package solution

import (
	"testing"
)

type Input struct {
	x float64
	n int
}

func TestProblem(t *testing.T) {
	tests := []struct {
		input  Input
		output float64
	}{
		{Input{2.0, 10}, 1024.00000},
		{Input{2.10000, 3}, 9.26100},
		{Input{2.00000, -2}, 0.25000},
	}
	for index, test := range tests {
		output := myPow(test.input.x, test.input.n)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("index: %v, expected: %v, output: %v", index, test.output, output)
		}
	}
}
