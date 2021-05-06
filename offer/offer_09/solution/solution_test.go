package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)

	var value int
	// case 1
	obj := Constructor()
	obj.AppendTail(3)
	value = obj.DeleteHead()
	assert.Equal(3, value)
	value = obj.DeleteHead()
	assert.Equal(-1, value)

	// case 2
	obj = Constructor()
	value = obj.DeleteHead()
	assert.Equal(-1, value)
	obj.AppendTail(5)
	obj.AppendTail(2)
	value = obj.DeleteHead()
	assert.Equal(5, value)
	value = obj.DeleteHead()
	assert.Equal(2, value)
}
