package linkedList_stack

import (
	"container/list"
	"sync"
)

type linkedListStack struct {
	stack *list.List // linkedList stack
	mu    sync.Mutex
}

// new linkedList stack
func NewStack() *linkedListStack {
	return &linkedListStack{
		stack: list.New(),
	}
}
