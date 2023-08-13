package algorithm

import (
	"github.com/khaosles/giz/g"
)

/*
   @File: numerical_comparator.go
   @Author: khaosles
   @Time: 2023/8/12 15:55
   @Desc:
*/

type NumericComparator[T g.Numeric] struct{}

func (nc NumericComparator[T]) Compare(lhs, rhs T) int {
	if lhs < rhs {
		return -1
	}
	if lhs > rhs {
		return 1
	}
	return 0
}
