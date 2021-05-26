package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input  []string
		k      int
		output []string
	}{
		{
			[]string{"i", "love", "leetcode", "i", "love", "coding"},
			2,
			[]string{"i", "love"},
		},
		{
			[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			4,
			[]string{"the", "is", "sunny", "day"},
		},
	}
	for index, test := range tests {
		output := topKFrequent(test.input, test.k)
		t.Logf("index: %v, input: %v, k: %v, output %v", index, test.input, test.k, output)
		assert.Equal(test.output, output)
	}
}
