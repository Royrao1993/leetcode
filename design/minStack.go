package design

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (s *MinStack) Push(i int) {
	s.stack = append(s.stack, i)
	top := s.minStack[len(s.stack)-1]
	s.minStack = append(s.minStack, min(top, i))
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
