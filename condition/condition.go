package condition

/*
   @File: condition.go
   @Author: khaosles
   @Time: 2023/8/13 09:30
   @Desc:
*/

import "reflect"

// Bool returns the truthy value of anything.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.
func Bool[T any](value T) bool {
	switch m := any(value).(type) {
	case interface{ Bool() bool }:
		return m.Bool()
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&value)
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() != 0
	default:
		is := rv.IsZero()
		return !is
	}
}

// And returns true if both a and b are truthy.
func And[T, U any](a T, b U) bool {
	return Bool(a) && Bool(b)
}

// Or returns false if neither a nor b is truthy.
func Or[T, U any](a T, b U) bool {
	return Bool(a) || Bool(b)
}

// Xor returns true if a or b but not both is truthy.
func Xor[T, U any](a T, b U) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA || valB) && valA != valB
}

// Nor returns true if neither a nor b is truthy.
func Nor[T, U any](a T, b U) bool {
	return !(Bool(a) || Bool(b))
}

// Xnor returns true if both a and b or neither a nor b are truthy.
func Xnor[T, U any](a T, b U) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA && valB) || (!valA && !valB)
}

// Nand returns false if both a and b are truthy.
func Nand[T, U any](a T, b U) bool {
	return !Bool(a) || !Bool(b)
}

// TernaryOperator checks the value of param `isTrue`, if true return ifValue else return elseValue.
func TernaryOperator[T, U any](isTrue T, ifValue U, elseValue U) U {
	if Bool(isTrue) {
		return ifValue
	} else {
		return elseValue
	}
}
