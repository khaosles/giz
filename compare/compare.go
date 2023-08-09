package compare

import (
	"reflect"
	"time"
)

/*
   @File: compare.go
   @Author: khaosles
   @Time: 2023/8/9 23:52
   @Desc:
*/

// operator type
const (
	equal          = "eq"
	lessThan       = "lt"
	greaterThan    = "gt"
	lessOrEqual    = "le"
	greaterOrEqual = "ge"
)

var (
	timeType  = reflect.TypeOf(time.Time{})
	bytesType = reflect.TypeOf([]byte{})
)

func Equal(left, right any) bool {
	return compareValue(equal, left, right)
}
