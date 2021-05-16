package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  []int
		k      int
		output int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for index, test := range tests {
		output := findKthLargest(test.input, test.k)
		t.Logf("index: %v, input: %v, k: %v, output %v", index, test.input, test.k, output)
		assert.Equal(test.output, output)
	}
}
