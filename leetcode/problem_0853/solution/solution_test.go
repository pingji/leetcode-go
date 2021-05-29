package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	targe    int
	position []int
	speed    []int
}

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  Input
		output int
	}{
		{
			Input{
				targe:    12,
				position: []int{10, 8, 0, 5, 3},
				speed:    []int{2, 4, 1, 1, 3},
			},
			3,
		},
	}
	for index, test := range tests {
		output := carFleet(test.input.targe, test.input.position, test.input.speed)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
