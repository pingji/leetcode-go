package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  [][]string
		output [][]string
	}{
		{
			[][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			[][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
	}
	for index, test := range tests {
		output := accountsMerge(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.ElementsMatch(test.output, output)
	}
}
