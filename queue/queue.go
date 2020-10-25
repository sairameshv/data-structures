// Package queue provides API implementation of few of the Queue operations.
package queue

type node struct {
	// data refers to the data that can be stored inside the node of a Queue
	data interface{}
	// reference refers to the pointer reference to the previous node of the queue
	reference *node
}

// Queue refers to the Queue (FIFO) data structure.
type Queue struct {
	nodes []node
	// front refers to the pointer reference to the front node present in the queue
	front *node
	// rear refers to the pointer reference to the rear node present in the queue
	rear *node
	// length refers to the number of nodes present in the queue
	length int
}

// NewQueue creates and returns a Queue variable
func NewQueue() *Queue {
	var queue = new(Queue)
	return queue
}

// Length returns the number of nodes present in the queue
func (q *Queue) Length() int {
	return q.length
}

// Dequeue removes a node from the queue and returns the popped out node's value
func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
	frontNode := q.front
	q.front = frontNode.reference
	q.nodes = q.nodes[:q.length-1]
	q.length--
	return frontNode.data
}

// Enqueue adds a new node to the existing queue
func (q *Queue) Enqueue(data interface{}) {
	newNode := node{data, q.rear}
	s.nodes = append(s.nodes, newNode)
	s.rear = &newNode
	s.length++
}

// Front returns the data value that is stored on front node of the queue
func (q *Queue) Front() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.front.data
}

// Rear returns the data value that is stored on rear node of the queue
func (q *Queue) Rear() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.front.data
}

// IsEmpty returns a true if the queue is empty and false otherwise.
func (q *Queue) IsEmpty() bool {
	if q.length == 0 {
		return true
	}
	return false
}
