package iterator

/*
   @File: operation.go
   @Author: khaosles
   @Time: 2023/8/13 10:50
   @Desc:
*/

// Map creates a new iterator which applies a function to all items of input iterator.
func Map[T any, U any](iter Iterator[T], iteratee func(item T) U) Iterator[U] {
	return &mapIterator[T, U]{
		iter:     iter,
		iteratee: iteratee,
	}
}

type mapIterator[T any, U any] struct {
	iter     Iterator[T]
	iteratee func(T) U
}

func (mr *mapIterator[T, U]) HasNext() bool {
	return mr.iter.HasNext()
}

func (mr *mapIterator[T, U]) Next() (U, bool) {
	var zero U
	item, ok := mr.iter.Next()
	if !ok {
		return zero, false
	}
	return mr.iteratee(item), true
}

// Filter creates a new iterator that returns only the items that pass specified predicate function.
func Filter[T any](iter Iterator[T], predicateFunc func(item T) bool) Iterator[T] {
	return &filterIterator[T]{iter: iter, predicateFunc: predicateFunc}
}

type filterIterator[T any] struct {
	iter          Iterator[T]
	predicateFunc func(T) bool
}

func (fr *filterIterator[T]) Next() (T, bool) {
	for item, ok := fr.iter.Next(); ok; item, ok = fr.iter.Next() {
		if fr.predicateFunc(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func (fr *filterIterator[T]) HasNext() bool {
	return fr.iter.HasNext()
}

// Join creates an iterator that join all elements of iters[0], then all elements of iters[1] and so on.
func Join[T any](iters ...Iterator[T]) Iterator[T] {
	return &joinIterator[T]{
		iters: iters,
	}
}

type joinIterator[T any] struct {
	iters []Iterator[T]
}

func (iter *joinIterator[T]) Next() (T, bool) {
	for len(iter.iters) > 0 {
		item, ok := iter.iters[0].Next()
		if ok {
			return item, true
		}
		iter.iters = iter.iters[1:]
	}
	var zero T
	return zero, false
}

func (iter *joinIterator[T]) HasNext() bool {
	if len(iter.iters) == 0 {
		return false
	}
	if len(iter.iters) == 1 {
		return iter.iters[0].HasNext()
	}

	result := iter.iters[0].HasNext()

	for i := 1; i < len(iter.iters); i++ {
		it := iter.iters[i]
		hasNext := it.HasNext()
		result = result || hasNext
	}

	return result
}

// Reduce reduces iter to a single value using the reduction function reducer
func Reduce[T any, U any](iter Iterator[T], initial U, reducer func(U, T) U) U {
	acc := initial

	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		acc = reducer(acc, item)
	}

	return acc
}

func Take[T any](it Iterator[T], num int) Iterator[T] {
	return &takeIterator[T]{it: it, num: num}
}

type takeIterator[T any] struct {
	it  Iterator[T]
	num int
}

func (iter *takeIterator[T]) Next() (T, bool) {
	if iter.num <= 0 {
		var zero T
		return zero, false
	}
	item, ok := iter.it.Next()
	if ok {
		iter.num--
	}
	return item, ok
}

func (iter *takeIterator[T]) HasNext() bool {
	return iter.num > 0
}
