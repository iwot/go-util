package util

import (
	//"fmt"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestStackString(t *testing.T) {
	stack := new(Stack)
	stack.Push("aaa")
	assert.Equal(t, 1, stack.Len(), "they should be equal")
	stack.Push("bbb")
	assert.Equal(t, 2, stack.Len(), "they should be equal")

	i := 0
	for elm := range stack.Iter() {
		i++
		switch i {
		case 1:
			assert.Equal(t, "bbb", elm.value, "they should be equal")
		case 2:
			assert.Equal(t, "aaa", elm.value, "they should be equal")
		default:
			assert.True(t, false, "out range error")
		}
	}

	stack.Push("cccc")
	assert.Equal(t, 3, stack.Len(), "they should be equal")

	i = 0
	for elm := range stack.Iter() {
		i++
		switch i {
		case 1:
			assert.Equal(t, "cccc", elm.value, "they should be equal")
		case 2:
			assert.Equal(t, "bbb", elm.value, "they should be equal")
		case 3:
			assert.Equal(t, "aaa", elm.value, "they should be equal")
		default:
			assert.True(t, false, "out range error")
		}
	}

	value, size := stack.Pop()
	assert.Equal(t, "cccc", value, "they should be equal")
	assert.Equal(t, 2, size, "they should be equal")
	assert.Equal(t, 2, stack.Len(), "they should be equal")

}

func TestStackNumber(t *testing.T) {
	stack := new(Stack)
	stack.Push(2)
	assert.Equal(t, 1, stack.Len(), "they should be equal")
	stack.Push(3)
	assert.Equal(t, 2, stack.Len(), "they should be equal")

	i := 0
	for elm := range stack.Iter() {
		i++
		switch i {
		case 1:
			assert.Equal(t, 3, elm.value, "they should be equal")
		case 2:
			assert.Equal(t, 2, elm.value, "they should be equal")
		default:
			assert.True(t, false, "out range error")
		}
	}

	stack.Push(4)
	assert.Equal(t, 3, stack.Len(), "they should be equal")

	i = 0
	for elm := range stack.Iter() {
		i++
		switch i {
		case 1:
			assert.Equal(t, 4, elm.value, "they should be equal")
		case 2:
			assert.Equal(t, 3, elm.value, "they should be equal")
		case 3:
			assert.Equal(t, 2, elm.value, "they should be equal")
		default:
			assert.True(t, false, "out range error")
		}
	}

	value, size := stack.Pop()
	assert.Equal(t, 4, value, "they should be equal")
	assert.Equal(t, 2, size, "they should be equal")
	assert.Equal(t, 2, stack.Len(), "they should be equal")

}
