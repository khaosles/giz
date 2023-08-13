package function

import (
	"fmt"
	"reflect"
	"time"
)

/*
   @File: function.go
   @Author: khaosles
   @Time: 2023/8/13 10:51
   @Desc:
*/

// After creates a function that invokes func once it's called n or more times.
func After(n int, fn any) func(args ...any) []reflect.Value {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	return func(args ...any) []reflect.Value {
		n--
		if n < 1 {
			return unsafeInvokeFunc(fn, args...)
		}
		return nil
	}
}

// Before creates a function that invokes func once it's called less than n times.
func Before(n int, fn any) func(args ...any) []reflect.Value {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)
	var result []reflect.Value
	return func(args ...any) []reflect.Value {
		if n > 0 {
			result = unsafeInvokeFunc(fn, args...)
		}
		if n <= 0 {
			fn = nil
		}
		n--
		return result
	}
}

// CurryFn is for make curry function
type CurryFn[T any] func(...T) T

// New make a curry function for specific value.
func (cf CurryFn[T]) New(val T) func(...T) T {
	return func(vals ...T) T {
		args := append([]T{val}, vals...)
		return cf(args...)
	}
}

// Compose  the functions from right to left.
func Compose[T any](fnList ...func(...T) T) func(...T) T {
	return func(args ...T) T {
		firstFn := fnList[0]
		restFns := fnList[1:]

		if len(fnList) == 1 {
			return firstFn(args...)
		}

		fn := Compose(restFns...)
		arg := fn(args...)

		return firstFn(arg)
	}
}

// Delay make the function execution after delayed time.
func Delay(delay time.Duration, fn any, args ...any) {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	time.Sleep(delay)
	unsafeInvokeFunc(fn, args...)
}

// Debounced creates a debounced function that delays invoking fn until after wait duration have elapsed since the last time the debounced function was invoked.
func Debounced(fn func(), duration time.Duration) func() {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	timer := time.NewTimer(duration)
	timer.Stop()

	go func() {
		for {
			<-timer.C
			go fn()
		}
	}()

	return func() { timer.Reset(duration) }
}

// Schedule invoke function every duration time, util close the returned bool channel.
func Schedule(d time.Duration, fn any, args ...any) chan bool {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	quit := make(chan bool)
	go func() {
		for {
			unsafeInvokeFunc(fn, args...)
			select {
			case <-time.After(d):
			case <-quit:
				return
			}
		}
	}()

	return quit
}

// Pipeline takes a list of functions and returns a function whose param will be passed into
// the functions one by one.
func Pipeline[T any](funcs ...func(T) T) func(T) T {
	return func(arg T) (result T) {
		result = arg
		for _, fn := range funcs {
			result = fn(result)
		}
		return
	}
}

func unsafeInvokeFunc(fn any, args ...any) []reflect.Value {
	fv := reflect.ValueOf(fn)
	params := make([]reflect.Value, len(args))
	for i, item := range args {
		params[i] = reflect.ValueOf(item)
	}
	return fv.Call(params)
}

func mustBeFunction(function any) {
	v := reflect.ValueOf(function)
	if v.Kind() != reflect.Func {
		panic(fmt.Sprintf("Invalid function type, value of type %T", function))
	}
}
