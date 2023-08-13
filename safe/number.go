package safe

import (
	"sync"

	"github.com/khaosles/giz/g"
)

/*
   @File: number.go
   @Author: khaosles
   @Time: 2023/5/24 17:21
   @Desc: 线程安全的数字
*/

type Number[T g.Numeric] struct {
	value T
	mutex sync.RWMutex
}

func NewNumber[T g.Numeric](i T) *Number[T] {
	return &Number[T]{value: i}
}

func (n *Number[T]) Inc() {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value++
}

func (n *Number[T]) Dec() {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value--
}

func (n *Number[T]) Add(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = n.value + val
}

func (n *Number[T]) Sub(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = n.value - val
}

func (n *Number[T]) Multiply(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = n.value * val
}

func (n *Number[T]) Divide(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = n.value / val
}

func (n *Number[T]) Get() T {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	return n.value
}

func (n *Number[T]) Set(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = val
}
