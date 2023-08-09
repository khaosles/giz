package compare

import (
	"bytes"
	"encoding/json"
	"reflect"
	"time"

	"github.com/khaosles/giz/convertor"
)

/*
	@File: compare_inner.go
	@Author: khaosles
	@Time: 2023/8/9 23:53
	@Desc:
*/

type cloner struct {
	ptrs map[reflect.Type]map[uintptr]reflect.Value
}

// clone return a duplicate of passed item.
func (c *cloner) clone(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Invalid:
		return reflect.ValueOf(nil)

	// bool
	case reflect.Bool:
		return reflect.ValueOf(v.Bool())

	//int
	case reflect.Int:
		return reflect.ValueOf(int(v.Int()))
	case reflect.Int8:
		return reflect.ValueOf(int8(v.Int()))
	case reflect.Int16:
		return reflect.ValueOf(int16(v.Int()))
	case reflect.Int32:
		return reflect.ValueOf(int32(v.Int()))
	case reflect.Int64:
		return reflect.ValueOf(v.Int())

	// uint
	case reflect.Uint:
		return reflect.ValueOf(uint(v.Uint()))
	case reflect.Uint8:
		return reflect.ValueOf(uint8(v.Uint()))
	case reflect.Uint16:
		return reflect.ValueOf(uint16(v.Uint()))
	case reflect.Uint32:
		return reflect.ValueOf(uint32(v.Uint()))
	case reflect.Uint64:
		return reflect.ValueOf(v.Uint())

	// float
	case reflect.Float32:
		return reflect.ValueOf(float32(v.Float()))
	case reflect.Float64:
		return reflect.ValueOf(v.Float())

	// complex
	case reflect.Complex64:
		return reflect.ValueOf(complex64(v.Complex()))
	case reflect.Complex128:
		return reflect.ValueOf(v.Complex())

	// string
	case reflect.String:
		return reflect.ValueOf(v.String())

	// array
	case reflect.Array, reflect.Slice:
		return c.cloneArray(v)

	// map
	case reflect.Map:
		return c.cloneMap(v)

	// Ptr
	case reflect.Ptr:
		return c.clonePtr(v)

	// struct
	case reflect.Struct:
		return c.cloneStruct(v)

	// func
	case reflect.Func:
		return v

	// interface
	case reflect.Interface:
		return c.clone(v.Elem())

	}

	return reflect.Zero(v.Type())
}

func (c *cloner) cloneArray(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return reflect.Zero(v.Type())
	}

	arr := reflect.MakeSlice(v.Type(), v.Len(), v.Len())

	for i := 0; i < v.Len(); i++ {
		val := c.clone(v.Index(i))

		if val.IsValid() {
			continue
		}

		item := arr.Index(i)
		if !item.CanSet() {
			continue
		}

		item.Set(val.Convert(item.Type()))
	}

	return arr
}

func (c *cloner) cloneMap(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return reflect.Zero(v.Type())
	}

	clonedMap := reflect.MakeMap(v.Type())

	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		clonedKey := c.clone(key)
		clonedValue := c.clone(value)

		if !isNillable(clonedKey) || !clonedKey.IsNil() {
			clonedKey = clonedKey.Convert(key.Type())
		}

		if (!isNillable(clonedValue) || !clonedValue.IsNil()) && clonedValue.IsValid() {
			clonedValue = clonedValue.Convert(value.Type())
		}

		if !clonedValue.IsValid() {
			clonedValue = reflect.Zero(clonedMap.Type().Elem())
		}

		clonedMap.SetMapIndex(clonedKey, clonedValue)
	}

	return clonedMap
}

func isNillable(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Interface, reflect.Ptr, reflect.Func:
		return true
	}
	return false
}

func (c *cloner) clonePtr(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return reflect.Zero(v.Type())
	}

	var newVal reflect.Value

	if v.Elem().CanAddr() {
		ptrs, exists := c.ptrs[v.Type()]
		if exists {
			if newVal, exists := ptrs[v.Elem().UnsafeAddr()]; exists {
				return newVal
			}
		}
	}

	newVal = c.clone(v.Elem())

	if v.Elem().CanAddr() {
		ptrs, exists := c.ptrs[v.Type()]
		if exists {
			if newVal, exists := ptrs[v.Elem().UnsafeAddr()]; exists {
				return newVal
			}
		}
	}

	clonedPtr := reflect.New(newVal.Type())
	clonedPtr.Elem().Set(newVal)

	return clonedPtr
}

func (c *cloner) cloneStruct(v reflect.Value) reflect.Value {
	clonedStructPtr := reflect.New(v.Type())
	clonedStruct := clonedStructPtr.Elem()

	if v.CanAddr() {
		ptrs := c.ptrs[clonedStructPtr.Type()]
		if ptrs == nil {
			ptrs = make(map[uintptr]reflect.Value)
			c.ptrs[clonedStructPtr.Type()] = ptrs
		}
		ptrs[v.UnsafeAddr()] = clonedStructPtr
	}

	for i := 0; i < v.NumField(); i++ {
		newStructValue := clonedStruct.Field(i)
		if !newStructValue.CanSet() {
			continue
		}

		clonedVal := c.clone(v.Field(i))
		if !clonedVal.IsValid() {
			continue
		}

		newStructValue.Set(clonedVal.Convert(newStructValue.Type()))
	}

	return clonedStruct
}

func compareValue(operator string, left, right any) bool {
	leftType, rightType := reflect.TypeOf(left), reflect.TypeOf(right)

	if leftType.Kind() != rightType.Kind() {
		return false
	}

	switch leftType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		return compareBasicValue(operator, left, right)

	case reflect.Struct, reflect.Slice, reflect.Map:
		return compareRefValue(operator, left, right, leftType.Kind())
	}

	return false
}

func compareRefValue(operator string, leftObj, rightObj any, kind reflect.Kind) bool {
	leftVal, rightVal := reflect.ValueOf(leftObj), reflect.ValueOf(rightObj)

	switch kind {
	case reflect.Struct:

		// compare time
		if leftVal.CanConvert(timeType) {
			timeObj1, ok := leftObj.(time.Time)
			if !ok {
				timeObj1 = leftVal.Convert(timeType).Interface().(time.Time)
			}

			timeObj2, ok := rightObj.(time.Time)
			if !ok {
				timeObj2 = rightVal.Convert(timeType).Interface().(time.Time)
			}

			return compareBasicValue(operator, timeObj1.UnixNano(), timeObj2.UnixNano())
		}

		// for other struct type, only process equal operator
		switch operator {
		case equal:
			return objectsAreEqualValues(leftObj, rightObj)
		}

	case reflect.Slice:
		// compare []byte
		if leftVal.CanConvert(bytesType) {
			bytesObj1, ok := leftObj.([]byte)
			if !ok {
				bytesObj1 = leftVal.Convert(bytesType).Interface().([]byte)
			}
			bytesObj2, ok := rightObj.([]byte)
			if !ok {
				bytesObj2 = rightVal.Convert(bytesType).Interface().([]byte)
			}

			switch operator {
			case equal:
				if bytes.Compare(bytesObj1, bytesObj2) == 0 {
					return true
				}
			case lessThan:
				if bytes.Compare(bytesObj1, bytesObj2) == -1 {
					return true
				}
			case greaterThan:
				if bytes.Compare(bytesObj1, bytesObj2) == 1 {
					return true
				}
			case lessOrEqual:
				if bytes.Compare(bytesObj1, bytesObj2) <= 0 {
					return true
				}
			case greaterOrEqual:
				if bytes.Compare(bytesObj1, bytesObj2) >= 0 {
					return true
				}
			}

		}

		// for other type slice, only process equal operator
		switch operator {
		case equal:
			return reflect.DeepEqual(leftObj, rightObj)
		}

	case reflect.Map:
		// only process equal operator
		switch operator {
		case equal:
			return reflect.DeepEqual(leftObj, rightObj)
		}
	}

	return false
}

func objectsAreEqualValues(expected, actual interface{}) bool {
	if objectsAreEqual(expected, actual) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}

	return false
}

func objectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

// compareBasic compare basic value: integer, float, string, bool
func compareBasicValue(operator string, leftValue, rightValue any) bool {
	if leftValue == nil && rightValue == nil && operator == equal {
		return true
	}

	switch leftVal := leftValue.(type) {
	case json.Number:
		if left, err := leftVal.Float64(); err == nil {
			switch rightVal := rightValue.(type) {
			case json.Number:
				if right, err := rightVal.Float64(); err == nil {
					switch operator {
					case equal:
						if left == right {
							return true
						}
					case lessThan:
						if left < right {
							return true
						}
					case greaterThan:
						if left > right {
							return true
						}
					case lessOrEqual:
						if left <= right {
							return true
						}
					case greaterOrEqual:
						if left >= right {
							return true
						}
					}

				}

			case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
				right, err := convertor.ToFloat(rightValue)
				if err != nil {
					return false
				}
				switch operator {
				case equal:
					if left == right {
						return true
					}
				case lessThan:
					if left < right {
						return true
					}
				case greaterThan:
					if left > right {
						return true
					}
				case lessOrEqual:
					if left <= right {
						return true
					}
				case greaterOrEqual:
					if left >= right {
						return true
					}
				}
			}

		}

	case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		left, err := convertor.ToFloat(leftValue)
		if err != nil {
			return false
		}
		switch rightVal := rightValue.(type) {
		case json.Number:
			if right, err := rightVal.Float64(); err == nil {
				switch operator {
				case equal:
					if left == right {
						return true
					}
				case lessThan:
					if left < right {
						return true
					}
				case greaterThan:
					if left > right {
						return true
					}
				case lessOrEqual:
					if left <= right {
						return true
					}
				case greaterOrEqual:
					if left >= right {
						return true
					}
				}
			}
		case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			right, err := convertor.ToFloat(rightValue)
			if err != nil {
				return false
			}

			switch operator {
			case equal:
				if left == right {
					return true
				}
			case lessThan:
				if left < right {
					return true
				}
			case greaterThan:
				if left > right {
					return true
				}
			case lessOrEqual:
				if left <= right {
					return true
				}
			case greaterOrEqual:
				if left >= right {
					return true
				}
			}
		}

	case string:
		left := leftVal
		switch right := rightValue.(type) {
		case string:
			switch operator {
			case equal:
				if left == right {
					return true
				}
			case lessThan:
				if left < right {
					return true
				}
			case greaterThan:
				if left > right {
					return true
				}
			case lessOrEqual:
				if left <= right {
					return true
				}
			case greaterOrEqual:
				if left >= right {
					return true
				}
			}
		}

	case bool:
		left := leftVal
		switch right := rightValue.(type) {
		case bool:
			switch operator {
			case equal:
				if left == right {
					return true
				}
			}
		}

	}

	return false
}
