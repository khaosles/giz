package pointer

import "reflect"

/*
   @File: pointer.go
   @Author: khaosles
   @Time: 2023/8/9 23:59
   @Desc:
*/

// Of returns a pointer to the value `v`.
func Of[T any](v T) *T {
	return &v
}

// Unwrap returns the value from the pointer.
func Unwrap[T any](p *T) T {
	return *p
}

// UnwarpOr returns the value from the pointer or fallback if the pointer is nil.
func UnwarpOr[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}
	return *p
}

// UnwarpOrDefault returns the value from the pointer or the default value if the pointer is nil.
func UnwarpOrDefault[T any](p *T) T {
	var v T

	if p == nil {
		return v
	}
	return *p
}

// ExtractPointer returns the underlying value by the given interface type
func ExtractPointer(value any) any {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if t.Kind() != reflect.Pointer {
		return value
	}
	return ExtractPointer(v.Elem().Interface())
}
