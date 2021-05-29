package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  []int
		output string
	}{
		{[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
		{[]int{1}, "1"},
		{[]int{10}, "10"},
	}
	for index, test := range tests {
		output := largestNumber(test.input)
		t.Logf("index: %v, input: %v, output %v", index, test.input, output)
		assert.Equal(test.output, output)
	}
}
