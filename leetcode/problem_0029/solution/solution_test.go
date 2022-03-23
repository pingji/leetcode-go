package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		nums   int
		val    int
		output int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{1, 1, 1},
		{-1, 1, -1},
		{-2147483648, -1, 2147483647},
	}
	for index, test := range tests {
		t.Logf("index: %v, nums: %v, val %v", index, test.nums, test.val)
		output := divide(test.nums, test.val)
		t.Logf("index: %v, output %v", index, output)
		assert.Equal(test.output, output)
	}
}
