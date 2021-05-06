package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)

	var value int
	// case 1
	myQueue := Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	value = myQueue.Peek()
	assert.Equal(1, value)
	value = myQueue.Pop()
	assert.Equal(1, value)

	empty := myQueue.Empty()
	assert.False(empty)
}
