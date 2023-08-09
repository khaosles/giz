package constraints

/*
   @File: comparator.go
   @Author: khaosles
   @Time: 2023/8/9 23:50
   @Desc:
*/

type Comparator interface {
	Compare(lhs, rhs any) int
}
