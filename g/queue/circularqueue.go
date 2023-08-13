package queue

import (
	"errors"
	"fmt"
	"reflect"
)

/*
   @File: circularqueue.go
   @Author: khaosles
   @Time: 2023/8/13 15:58
   @Desc:
*/

// CircularQueue implements circular queue with slice,
// last index of CircularQueue don't contain value, so acturl capacity is capacity - 1
type CircularQueue[T any] struct {
	data     []T
	front    int
	rear     int
	capacity int
}

// NewCircularQueue return a empty CircularQueue pointer
func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	data := make([]T, capacity)
	return &CircularQueue[T]{data: data, front: 0, rear: 0, capacity: capacity}
}

// Data return slice of queue data
func (q *CircularQueue[T]) Data() []T {
	data := []T{}

	front := q.front
	rear := q.rear
	if front <= rear {
		return q.data[front:rear]
	}

	data = append(data, q.data[front:]...)
	data = append(data, q.data[0:rear]...)

	return data
}

// Size return number of elements in circular queue
func (q *CircularQueue[T]) Size() int {
	if q.capacity == 0 {
		return 0
	}
	return (q.rear - q.front + q.capacity) % q.capacity
}

// IsEmpty checks if queue is empty or not
func (q *CircularQueue[T]) IsEmpty() bool {
	return q.front == q.rear
}

// IsFull checks if queue is full or not
func (q *CircularQueue[T]) IsFull() bool {
	return (q.rear+1)%q.capacity == q.front
}

// Front return front value of queue
func (q *CircularQueue[T]) Front() T {
	return q.data[q.front]
}

// Back return back value of queue
func (q *CircularQueue[T]) Back() T {
	if q.rear-1 >= 0 {
		return q.data[q.rear-1]
	}
	return q.data[q.capacity-1]
}

// Enqueue put element into queue
func (q *CircularQueue[T]) Enqueue(value T) error {
	if q.IsFull() {
		return errors.New("queue is full!")
	}

	q.data[q.rear] = value
	q.rear = (q.rear + 1) % q.capacity

	return nil
}

// Dequeue remove head element of queue and return it, if queue is empty, return nil and error
func (q *CircularQueue[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	headItem := q.data[q.front]
	var t T
	q.data[q.front] = t
	q.front = (q.front + 1) % q.capacity

	return &headItem, nil
}

// Clear the queue data
func (q *CircularQueue[T]) Clear() {
	q.data = []T{}
	q.front = 0
	q.rear = 0
	q.capacity = 0
}

// Contain checks if the value is in queue or not
func (q *CircularQueue[T]) Contain(value T) bool {
	for _, v := range q.data {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Print queue data
func (q *CircularQueue[T]) Print() {
	fmt.Printf("%+v\n", q)
}
