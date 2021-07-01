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
		{[]int{1, 1, 2}, 5},
		{[]int{10, 10, 10}, 11},
		{[]int{}, 0},
	}
	for index, test := range tests {
		output := numRabbits(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
