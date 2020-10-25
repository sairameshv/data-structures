# queue
--
    import "data-structures/queue"

Package queue provides API implementation of few of the Queue operations.

## Usage

#### type Queue

```go
type Queue struct {
}
```

Queue refers to the Queue (FIFO) data structure.

#### func  NewQueue

```go
func NewQueue() *Queue
```
NewQueue creates and returns a Queue variable

#### func (*Queue) Dequeue

```go
func (q *Queue) Dequeue() interface{}
```
Dequeue removes a node from the queue and returns the popped out node's value

#### func (*Queue) Enqueue

```go
func (q *Queue) Enqueue(data interface{})
```
Enqueue adds a new node to the existing queue

#### func (*Queue) Front

```go
func (q *Queue) Front() interface{}
```
Front returns the data value that is stored on front node of the queue

#### func (*Queue) IsEmpty

```go
func (q *Queue) IsEmpty() bool
```
IsEmpty returns a true if the queue is empty and false otherwise.

#### func (*Queue) Length

```go
func (q *Queue) Length() int
```
Length returns the number of nodes present in the queue

#### func (*Queue) Rear

```go
func (q *Queue) Rear() interface{}
```
Rear returns the data value that is stored on rear node of the queue
