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
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
	}
	for index, test := range tests {
		output := wiggleMaxLength(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
