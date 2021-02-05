package array_stack

import "testing"

func TestArrayStack(t *testing.T) {
	stack := NewStack(10)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	var v interface{}
	var err error
	v, err = stack.Peek()
	t.Log(v.(int), err)
	v, err = stack.Pop()
	t.Log(v.(int), err)
	v, err = stack.Peek()
	t.Log(v.(int), err)
	t.Log(stack.IsEmpty())
	t.Log(stack.Search(1))
}
