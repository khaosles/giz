package safe

import (
	"sync"
)

/*
   @File: map.go
   @Author: khaosles
   @Time: 2023/6/20 10:23
   @Desc:
*/

type Map[K comparable, V any] struct {
	m     map[K]V
	mutex sync.Mutex
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: map[K]V{}}
}

func (s *Map[K, V]) Put(key K, value V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[key] = value
}

func (s *Map[K, V]) HasKey(key K) bool {
	_, ok := s.m[key]
	return ok
}

func (s *Map[K, V]) Get(key K) (V, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if !s.HasKey(key) {
		return *new(V), false
	}
	return s.m[key], true
}

func (s *Map[K, V]) Del(key K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.m, key)
}

func (s *Map[K, V]) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.m)
}

func (s *Map[K, V]) Value() map[K]V {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.m
}

func (s *Map[K, V]) Range(cb func(key K, value V)) {
	for k, v := range s.m {
		cb(k, v)
	}
}
