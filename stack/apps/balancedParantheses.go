// Package apps refers to few of the Stack applications.
package apps

import (
	stk "data-structures/stack"
	"log"
)

var (
	paranthesesMap = map[int32]int32{
		'(': ')',
		'{': '}',
		'[': ']',
	}
)

// isValidInput validates the provided input string and returns the boolean accordingly
func isValidInput(input string) bool {
	for _, v := range input {
		switch v {
		case '(', ')', '{', '}', '[', ']':
			continue
		default:
			return false
		}
	}
	return true
}

// IsBalancedParantheses validates if the provided input sequence of parantheses is balanced
// and returns a boolean accordingly
func IsBalancedParantheses(input string) bool {
	if !isValidInput(input) {
		log.Printf("Invalid Input: %v", input)
		return false
	}
	stack := stk.NewStack()
	for _, v := range input {
		switch v {
		case '(', '{', '[':
			stack.Push(v)
		default:
			value := stack.Pop()
			if value == nil {
				return false
			}
			item := value.(int32)
			if paranthesesMap[item] != v {
				return false
			}
			continue
		}
	}
	return true
}
