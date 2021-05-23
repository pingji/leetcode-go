package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	assert := assert.New(t)

	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)

	var val int
	val = obj.Add(3)
	assert.Equal(4, val)

	val = obj.Add(5)
	assert.Equal(5, val)

	val = obj.Add(10)
	assert.Equal(5, val)

	val = obj.Add(9)
	assert.Equal(8, val)

	val = obj.Add(4)
	assert.Equal(8, val)
}
