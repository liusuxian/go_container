package array_stack

import "testing"

func TestArrayStack(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)

	var v interface{}
	var err error

	v, err = stack.Peek()
	t.Log("111: ", v.(int), err, stack.topIndex)

	v, err = stack.Pop()
	t.Log("222: ", v.(int), err, stack.topIndex)
	t.Log("222: ", stack.IsEmpty(), stack.topIndex)
	t.Log("222: ", stack.Search(1), stack.topIndex)

	v, err = stack.Pop()
	t.Log("333: ", v.(int), err, stack.topIndex)
	t.Log("333: ", stack.IsEmpty(), stack.topIndex)

	stack.Push(3)
	v, err = stack.Peek()
	t.Log("444: ",v.(int), err, stack.topIndex)

	v, err = stack.Pop()
	t.Log("555: ",v.(int), err, stack.topIndex)

	stack.Push(4)
	t.Log("666: ",v.(int), err, stack.topIndex)
}
