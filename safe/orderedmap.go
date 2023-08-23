package safe

import (
	"fmt"
	"sync"
)

/*
   @File: orderedmap.go
   @Author: khaosles
   @Time: 2023/8/9 00:01
   @Desc:
*/

// OrderedMap represents an associative array or map abstract data type.
type OrderedMap[K comparable, V any] struct {
	// mu Mutex protects data structures below.
	mu sync.RWMutex

	// keys is the Set list of keys.
	keys []K

	// store is the Set underlying store of values.
	store map[K]V
}

// NewOrderedMap creates a new empty OrderedMap.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	m := &OrderedMap[K, V]{
		keys:  make([]K, 0),
		store: make(map[K]V),
	}
	return m
}

// Put adds items to the map.
//
// If a key is found in the map it replaces it value.
func (m *OrderedMap[K, V]) Put(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.store[key]; !ok {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get returns the value of a key from the OrderedMap.
func (m *OrderedMap[K, V]) Get(key K) (value V, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok = m.store[key]
	return value, ok
}

// Remove deletes a key-value pair from the OrderedMap.
//
// If a key is not found in the map it doesn't fails, just does nothing.
func (m *OrderedMap[K, V]) Remove(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check key exists
	if _, ok := m.store[key]; !ok {
		return
	}

	// Remove the value from the store
	delete(m.store, key)

	// Remove the key
	for i := range m.keys {
		if m.keys[i] == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

// Foreach iter map.
func (m *OrderedMap[K, V]) Foreach(callback func(key K, value V) bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range m.store {
		if !callback(k, v) {
			break
		}
	}
}

// Size return the map number of key-value pairs.
func (m *OrderedMap[K, V]) Size() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	return len(m.store)
}

// Empty return if the map in empty or not.
func (m *OrderedMap[K, V]) Empty() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return len(m.store) == 0
}

// Keys return the keys in the map in insertion order.
func (m *OrderedMap[K, V]) Keys() []K {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.keys
}

// Values return the values in the map in insertion order.
func (m *OrderedMap[K, V]) Values() []V {
	m.mu.Lock()
	defer m.mu.Unlock()

	values := make([]V, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

func (m *OrderedMap[K, V]) Data() map[K]V {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.store
}

// String implements Stringer interface.
//
// Prints the map string representation, a concatenated string of all its
// string representation values in insertion order.
func (m *OrderedMap[K, V]) String() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	var result string
	for i, key := range m.keys {
		result += fmt.Sprintf("\t%v:%v, ", m.keys[i], m.store[key])
	}

	return fmt.Sprintf("{\n%s\n}", result)
}
