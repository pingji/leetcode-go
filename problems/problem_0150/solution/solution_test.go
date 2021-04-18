package solution

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []string
		output int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for index, test := range tests {
		output := evalRPN(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		if output != test.output {
			t.Errorf("Error! expected: %v, output: %v", test.output, output)
		}
	}
}
