package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	g []int
	s []int
}

func TestProblem(t *testing.T) {
	tests := []struct {
		input  Input
		output int
	}{
		{
			Input{
				[]int{1, 2, 3},
				[]int{1, 1},
			},
			1,
		},
		{
			Input{
				[]int{1, 2},
				[]int{1, 2, 3},
			},
			2,
		},
	}
	assert := assert.New(t)
	for index, test := range tests {
		output := findContentChildren(test.input.g, test.input.s)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
