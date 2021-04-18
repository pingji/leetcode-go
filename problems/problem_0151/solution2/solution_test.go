package solution

import "testing"

func TestProblem(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
	}
	for index, test := range tests {
		output := reverseWords(test.input)
		t.Logf("index: %v, input: %q, output: %q", index, test.input, output)
		if output != test.output {
			t.Errorf("index: %v, expected: %q, output: %q", index, test.output, output)
		}
	}
}
