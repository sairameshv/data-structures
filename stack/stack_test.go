// Package stack provides API implementation of few of the Stack operations.
package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := new(Stack)
	tests := []struct {
		name string
		want *Stack
	}{
		// Testcase to validate the creation of new stack
		{"Testcase-1", stack},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Length(t *testing.T) {
	type fields struct {
		node   node
		top    *node
		length int
	}
	stack := new(fields)
	testStack := NewStack()
	testStack.Push("Random Data")
	newfields := new(fields)
	newfields.node = testStack.node
	newfields.top = testStack.top
	newfields.length = testStack.length
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// Testcase for an empty stack
		{"Testcase-2", *stack, 0},
		// Testcase for a non-empty stack
		{"Testcase-3", *newfields, newfields.length},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				node:   tt.fields.node,
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			if got := s.Length(); got != tt.want {
				t.Errorf("Stack.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		node   node
		top    *node
		length int
	}
	// creating a test stack
	testStack := NewStack()
	testStack.Push("Random Data")
	// creating new fields
	newfields := new(fields)
	newfields.node = testStack.node
	newfields.top = testStack.top
	newfields.length = testStack.length

	// creating an empty stack
	emptyStack := NewStack()
	// creating empty fields
	emptyfields := new(fields)

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// Testcase to validate the Pop functionality
		{"Testcase-4", *newfields, testStack.Pop()},
		// Testcase to validate the Pop of an empty stack
		{"Testcase-5", *emptyfields, emptyStack.Pop()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				node:   tt.fields.node,
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			if got := s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		node   node
		top    *node
		length int
	}
	type args struct {
		data interface{}
	}
	// creating a test stack
        testStack := NewStack()
        // creating new fields
        newfields := new(fields)
        newfields.node = testStack.node
        newfields.top = testStack.top
        newfields.length = testStack.length
	// Pushing the data to the stack
        testStack.Push("Random Data")

	tests := []struct {
		name   string
		fields fields
		args   args
		// want represents the length of the updated stack
		want   int
	}{
		{"Testcase-6", *newfields, args{"Random Data"}, testStack.Length()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				node:   tt.fields.node,
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			s.Push(tt.args.data)
			if got := s.Length(); !reflect.DeepEqual(got, tt.want) {
                                t.Errorf("Stack.Push() = %v, want %v", got, tt.want)
                        }
		})
	}
}

func TestStack_Top(t *testing.T) {
	type fields struct {
		node   node
		top    *node
		length int
	}
	// creating a test stack
        testStack := NewStack()
	testStack.Push("Random Data")
        // creating new fields
        newfields := new(fields)
        newfields.node = testStack.node
        newfields.top = testStack.top
        newfields.length = testStack.length
        // creating empty fields
        emptyfields := new(fields)

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// Testcase for an empty stack
		{"Testcase-7", *emptyfields, nil},
		// Testcase for a non empty stack
		{"Testcase-8", *newfields, testStack.Top()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				node:   tt.fields.node,
				top:    tt.fields.top,
				length: tt.fields.length,
			}
			if got := s.Top(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}
