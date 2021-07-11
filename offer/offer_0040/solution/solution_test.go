package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		output []int
	}{
		{[]int{3, 2, 1}, 2, []int{1, 2}},
		{[]int{0, 1, 2, 1}, 1, []int{0}},
	}
	assert := assert.New(t)
	for index, test := range tests {
		output := getLeastNumbers(test.input, test.k)
		t.Logf("index: %v, input: %v, k: %v, output: %v", index, test.input, test.k, output)
		assert.ElementsMatch(test.output, output)
	}
}
