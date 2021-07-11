package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	k      int
	prices []int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			Input{
				2,
				[]int{2, 4, 1},
			},
			2,
		},
		{
			Input{
				2,
				[]int{3, 2, 6, 5, 0, 3},
			},
			7,
		},
	}
	for index, test := range tests {
		output := maxProfit(test.input.k, test.input.prices)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
