package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)

	var value int
	// case 1
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	value = minStack.Min()
	assert.Equal(-3, value)
	minStack.Pop()
	value = minStack.Top()
	assert.Equal(0, value)
	value = minStack.Min()
	assert.Equal(-2, value)
}
