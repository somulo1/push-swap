package main

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{values: make([]int, 0)}
}

func (s *Stack) Push(value int) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return value, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	return s.values[len(s.values)-1], true
}

func (s *Stack) Size() int {
	return len(s.values)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) IsSorted() bool {
	if len(s.values) <= 1 {
		return true
	}
	// Stack is sorted if elements from top to bottom are in ascending order
	// Top of stack is at the end of the slice
	for i := len(s.values) - 1; i > 0; i-- {
		if s.values[i] > s.values[i-1] {
			return false
		}
	}
	return true
}

func (s *Stack) Clone() *Stack {
	clone := NewStack()
	clone.values = make([]int, len(s.values))
	copy(clone.values, s.values)
	return clone
}

func (s *Stack) GetValues() []int {
	result := make([]int, len(s.values))
	copy(result, s.values)
	return result
}

func (s *Stack) Top() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	return s.values[len(s.values)-1], true
}

func (s *Stack) Second() (int, bool) {
	if len(s.values) < 2 {
		return 0, false
	}
	return s.values[len(s.values)-2], true
}
