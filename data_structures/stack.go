package data_structures

import (
	"sync"
)

type Stack struct {
	items        []interface{}
	mutexChannel sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		items:        make([]interface{}, 0),
		mutexChannel: sync.Mutex{},
	}
}

func (s *Stack) Push(item interface{}) {
	s.mutexChannel.Lock()
	defer s.mutexChannel.Unlock()
	s.items = append(s.items, item)
}
func (s *Stack) Pop() interface{} {
	s.mutexChannel.Lock()
	defer s.mutexChannel.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
func (s *Stack) Size() int {
	s.mutexChannel.Lock()
	defer s.mutexChannel.Unlock()
	return len(s.items)
}
