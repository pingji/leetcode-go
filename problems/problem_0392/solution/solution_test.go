package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		s      string
		t      string
		output bool
	}{
		{"", "ahbgdc", true},
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"aaaaaa", "bbaaaa", false},
	}
	for index, test := range tests {
		output := isSubsequence(test.s, test.t)
		t.Logf("index: %v, s: %v, t: %v, output %v", index, test.s, test.t, output)
		assert.Equal(test.output, output)
	}
}
