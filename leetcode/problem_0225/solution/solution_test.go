package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)

	var value int
	// case 1
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	value = myStack.Top()
	assert.Equal(2, value)
	value = myStack.Pop()
	assert.Equal(2, value)
	empty := myStack.Empty()
	assert.False(empty)
}
