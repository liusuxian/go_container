package array_stack

import (
	"errors"
	"sync"
)

type ArrayStack struct {
	stack    []interface{} // array stack
	topIndex int           // stack top index
	mu       sync.Mutex
}

// new array stack
func NewStack(cap int) *ArrayStack {
	return &ArrayStack{
		stack:    make([]interface{}, 0, cap),
		topIndex: -1,
	}
}

func (as *ArrayStack) Push(v interface{}) interface{} {
	as.mu.Lock()
	defer as.mu.Unlock()

	if len(as.stack) == 0 && as.topIndex < 0 {
		as.stack = append(as.stack, v)
		as.topIndex = 0
	} else if len(as.stack) > as.topIndex+1 {
		as.topIndex++
		as.stack[as.topIndex] = v
	} else {
		as.stack = append(as.stack, v)
		as.topIndex++
	}
	return v
}

func (as *ArrayStack) Pop() (interface{}, error) {
	as.mu.Lock()
	defer as.mu.Unlock()

	if as.topIndex < 0 {
		return nil, errors.New("stack is empty")
	}

	v := as.stack[as.topIndex]
	as.stack[as.topIndex] = nil
	as.topIndex--
	return v, nil
}

func (as *ArrayStack) Peek() (interface{}, error) {
	as.mu.Lock()
	defer as.mu.Unlock()

	if as.topIndex < 0 {
		return nil, errors.New("stack is empty")
	}

	return as.stack[as.topIndex], nil
}

func (as *ArrayStack) IsEmpty() bool {
	as.mu.Lock()
	defer as.mu.Unlock()

	if as.topIndex < 0 {
		return true
	}

	return false
}

func (as *ArrayStack) Search(v interface{}) int {
	as.mu.Lock()
	defer as.mu.Unlock()

	for i, val := range as.stack {
		if val == v {
			return i
		}
	}

	return -1
}
