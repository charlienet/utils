package stack_test

import (
	"testing"

	"github.com/charlienet/utils/stack"
)

func TestStack(t *testing.T) {
	arrayStack := new(stack.ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")

	t.Log("size:", arrayStack.Size())
	t.Log("pop:", arrayStack.Pop())
	t.Log("pop:", arrayStack.Pop())
	t.Log("size:", arrayStack.Size())
	arrayStack.Push("drag")
	t.Log("pop:", arrayStack.Pop())
	arrayStack.Push("test")
	s := arrayStack.Pop().(string)
	t.Log(s)
}
