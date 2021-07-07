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
		{[]int{2, 3, 1, 0, 2, 5, 3}, 2},
	}
	for index, test := range tests {
		output := findRepeatNumber(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
