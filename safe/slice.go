package safe

import (
	"fmt"
	"sync"
)

/*
   @File: slice.go
   @Author: khaosles
   @Time: 2023/8/6 09:22
   @Desc:
*/

type Slice[T any] struct {
	v   []*T
	m   sync.Mutex
	len int
}

func NewSafeSlice[T any]() *Slice[T] {
	return &Slice[T]{len: 0}
}

// Add 添加元素
func (s *Slice[T]) Add(val *T) {
	s.m.Lock()
	defer s.m.Unlock()
	s.v = append(s.v, val)
	s.len++
}

// Pop 删除元素
func (s *Slice[T]) Pop(i int) error {
	s.m.Lock()
	defer s.m.Unlock()
	if i < s.len {
		return fmt.Errorf("index %d out of bounds for length %d", i, s.len)
	}
	copy(s.v[i:], s.v[i+1:])
	s.v = s.v[:s.len-1]
	s.len--
	return nil
}

// Get 获去元素
func (s *Slice[T]) Get(i int) (*T, error) {
	s.m.Lock()
	defer s.m.Unlock()
	if i < s.len {
		return nil, fmt.Errorf("index %d out of bounds for length %d", i, s.len)
	}
	return s.v[i], nil
}

// Iter 获得一个迭代体
func (s *Slice[T]) Iter() []*T {
	s.m.Lock()
	defer s.m.Unlock()
	return s.v
}

func (s *Slice[T]) Len() int {
	return s.len
}

// Clear 删除全部元素
func (s *Slice[T]) Clear() {
	s.m.Lock()
	defer s.m.Unlock()
	s.len = 0
	s.v = make([]*T, 0)
}
