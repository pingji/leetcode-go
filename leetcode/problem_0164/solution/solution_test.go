package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	nums []int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			input: Input{
				[]int{3, 6, 9, 1},
			},
			output: 3,
		},
		{
			input: Input{
				[]int{10},
			},
			output: 0,
		},
	}
	for index, test := range tests {
		output := maximumGap(test.input.nums)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
