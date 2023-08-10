package compare

/*
   @File: compare_t.go
   @Author: khaosles
   @Time: 2023/8/10 16:33
   @Desc:
*/

// Eq whether left equal right
func Eq[T comparable](left, right T) bool {
	return left == right
}
