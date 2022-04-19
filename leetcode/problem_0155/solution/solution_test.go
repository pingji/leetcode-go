package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)
	// case 1
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(-3, minStack.GetMin())
	minStack.Pop()
	assert.Equal(0, minStack.Top())
	assert.Equal(-2, minStack.GetMin())
}
