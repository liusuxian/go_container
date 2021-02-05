package array_stack

import (
	"errors"
	"sync"
)

type ArrayStack struct {
	stack []interface{}

	mu sync.Mutex
}

func NewStack(cap int) *ArrayStack {
	return &ArrayStack{
		stack: make([]interface{}, 0, cap),
	}
}

func (as *ArrayStack) Push(v interface{}) interface{} {
	as.mu.Lock()
	defer as.mu.Unlock()
	as.stack = append(as.stack, v)
	return v
}

func (as *ArrayStack) Pop() (interface{}, error) {
	as.mu.Lock()
	defer as.mu.Unlock()

	if len(as.stack) == 0 {
		return nil, errors.New("stack is empty")
	}

	v := as.stack[len(as.stack)-1]
	as.stack = as.stack[:len(as.stack)-1]
	return v, nil
}

func (as *ArrayStack) Peek() (interface{}, error) {
	as.mu.Lock()
	defer as.mu.Unlock()

	if len(as.stack) == 0 {
		return nil, errors.New("stack is empty")
	}

	return as.stack[len(as.stack)-1], nil
}

func (as *ArrayStack) IsEmpty() bool {
	as.mu.Lock()
	defer as.mu.Unlock()

	if len(as.stack) == 0 {
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
