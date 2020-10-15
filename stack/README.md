# stack
--
    import "data-structures/stack"

Package stack provides API implementation of few of the Stack operations.

## Usage

#### type Stack

```go
type Stack struct {
}
```

Stack refers to the Stack (LIFO/FILO) data structure.

#### func  NewStack

```go
func NewStack() *Stack
```
NewStack creates and returns a Stack variable

#### func (*Stack) IsEmpty

```go
func (s *Stack) IsEmpty() bool
```
IsEmpty returns a true if the stack is empty and false otherwise.

#### func (*Stack) Length

```go
func (s *Stack) Length() int
```
Length returns the number of nodes present in the stack

#### func (*Stack) Pop

```go
func (s *Stack) Pop() interface{}
```
Pop deletes the top node from the stack and returns the popped out node's value

#### func (*Stack) Push

```go
func (s *Stack) Push(data interface{})
```
Push adds a new node to the existing stack

#### func (*Stack) Top

```go
func (s *Stack) Top() interface{}
```
Top returns the data value that is stored on top node of the stack
