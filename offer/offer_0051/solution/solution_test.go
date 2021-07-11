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
		{[]int{7, 5, 6, 4}, 5},
		{[]int{1, 3, 2, 3, 1}, 4},
		{[]int{}, 0},
	}
	for index, test := range tests {
		output := reversePairs(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
