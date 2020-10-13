// Package stack provides API implementation of few of the Stack operations.
package stack

type node struct {
	// data refers to the data that can be stored inside the node of a Stack
	data interface{}
	// reference refers to the pointer reference to the previous node of the stack
	reference *node
}

// Stack refers to the Stack (LIFO/FILO) data structure.
type Stack struct {
	node
	// previous refers to the pointer reference to the previous node present in the stack
	top *node
	// length refers to the number of nodes present in the stack
	length int
}

// NewStack creates and returns a Stack variable
func NewStack() *Stack {
	var stack = new(Stack)
	return stack
}

// Length returns the number of nodes present in the stack
func(s *Stack) Length() int {
	return s.length
}

// Pop deletes the top node from the stack and returns the popped out node's value
func(s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	topNode := s.top
	s.top = s.node.reference
	s.length--
	return topNode.data
}

// Push adds a new node to the existing stack
func(s *Stack) Push(data interface{}) {
	s.node = node{data, s.top}
	s.top = &s.node
	s.length++
}

// Top refers to the data value that is stored on top of the stack
func(s * Stack) Top() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.data
}
