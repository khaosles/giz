package constraints

/*
   @File: comparator.go
   @Author: khaosles
   @Time: 2023/8/9 23:50
   @Desc:
*/

type Comparator[T any] interface {
	// Compare .
	// if lhs < rhs  return < 0
	// if lhs = rhs  return = 0
	// if lhs > rhs  return > 0
	Compare(lhs, rhs T) int
}
