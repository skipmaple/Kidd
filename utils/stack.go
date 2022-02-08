package utils

import (
	"fmt"
	"sync"
)

type Item interface{}

type Stack struct {
	Items []Item
	lock  sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		Items: []Item{},
	}
}

func (s *Stack) Push(t Item) {
	s.lock.Lock()
	s.lock.Unlock()
	s.Items = append(s.Items, t)
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) > 0 {
		return false
	}
	return true
}

func (s *Stack) Pop() Item {
	s.lock.Lock()
	s.lock.Unlock()
	if s.IsEmpty() {
		return nil
	}
	res := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return res
}

func (s *Stack) TopValue() Item {
	s.lock.Lock()
	s.lock.Unlock()
	if s.IsEmpty() {
		return nil
	}
	return s.Items[len(s.Items)-1]
}

func (s *Stack) Print() {
	fmt.Println(s.Items)
}
