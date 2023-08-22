package safe

import (
	"fmt"
	"sync"

	"github.com/bytedance/sonic"
)

/*
   @File: orderedset.go
   @Author: khaosles
   @Time: 2023/8/9 00:01
   @Desc:
*/

// OrderedSet represents a dynamic, insertion-ordered, set abstract data type.
type OrderedSet[T any] struct {
	// currentIndex keeps track of the keys of the underlying store.
	currentIndex int

	// mu Mutex protects data structures below.
	mu sync.Mutex

	// index is the Set list of keys.
	index map[any]int

	// store is the Set underlying store of values.
	store *OrderedMap[int, T]
}

// NewOrderedSet creates a new empty OrderedSet.
func NewOrderedSet[T comparable]() *OrderedSet[T] {
	orderedset := &OrderedSet[T]{
		index: make(map[any]int),
		store: NewOrderedMap[int, T](),
	}

	return orderedset
}

// Add adds items to the set.
//
// If an item is found in the set it replaces it.
func (s *OrderedSet[T]) Add(items ...T) {

	for _, item := range items {
		if _, ok := s.index[item]; ok {
			continue
		}

		s.put(item)
	}
}

// Remove deletes items from the set.
//
// If an item is not found in the set it doesn't fails, just does nothing.
func (s *OrderedSet[T]) Remove(items ...T) {

	for _, item := range items {
		index, ok := s.index[item]
		if !ok {
			return
		}

		s.remove(index, item)
	}
}

// Contains return if set contains the specified items or not.
func (s *OrderedSet[T]) Contains(items ...T) bool {

	for _, item := range items {
		if _, ok := s.index[item]; !ok {
			return false
		}
	}
	return true
}

// Empty return if the set in empty or not.
func (s *OrderedSet[T]) Empty() bool {

	return s.store.Empty()
}

// Values return the set values in insertion order.
func (s *OrderedSet[T]) Values() []T {
	return s.store.Values()
}

// Size return the set number of elements.
func (s *OrderedSet[T]) Size() int {
	return s.store.Size()
}

// String implements Stringer interface.
//
// Prints the set string representation, a concatenated string of all its
// string representation values in insertion order.
func (s *OrderedSet[T]) String() string {
	return fmt.Sprintf("%s", s.Values())
}

// Foreach iter set.
func (s *OrderedSet[T]) Foreach(callback func(i int, val T) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, val := range s.store.Values() {
		if !callback(i, val) {
			break
		}
	}
}

// FromList converts a list to an OrderedSet.
func (s *OrderedSet[T]) FromList(list []T) {
	for _, item := range list {
		s.Add(item)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (s *OrderedSet[T]) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(s.Values())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *OrderedSet[T]) UnmarshalJSON(data []byte) error {
	var list []T
	if err := sonic.Unmarshal(data, &list); err != nil {
		return err
	}
	s.FromList(list)
	return nil
}

// Put adds a single item into the set
func (s *OrderedSet[T]) put(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store.Put(s.currentIndex, item)
	s.index[item] = s.currentIndex
	s.currentIndex++
}

// remove deletes a single item from the test given its index
func (s *OrderedSet[T]) remove(index int, item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store.Remove(index)
	delete(s.index, item)
}
