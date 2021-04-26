package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}
	for index, test := range tests {
		t.Logf("index: %v, input: %v", index, test.input)
		moveZeroes(test.input)
		t.Logf("index: %v, output %v", index, test.input)
		assert.Equal(test.output, test.input)
	}
}
