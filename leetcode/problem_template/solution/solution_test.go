package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3}, 0},
		{[]int{1, 2, 3}, 0},
		{[]int{1, 2, 3}, 0},
	}
	for index, test := range tests {
		output := foo(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
