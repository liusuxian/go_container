package linkedList_stack

import (
	"container/list"
	"sync"
)

type LinkedListStack struct {
	stack *list.List // linkedList stack
	mu    sync.Mutex
}
