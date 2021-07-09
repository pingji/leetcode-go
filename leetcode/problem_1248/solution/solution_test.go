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
		{[]int{1, 1, 2, 1, 1}, 3, 2},
		{[]int{2, 4, 6}, 1, 0},
		{[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2, 16},
	}
	for index, test := range tests {
		output := numberOfSubarrays(test.input, test.k)
		t.Logf("index: %v, input: %v, k: %v, output %v", index, test.input, test.k, output)
		assert.Equal(test.output, output)
	}
}
