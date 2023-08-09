package structs

import "fmt"

/*
   @File: struct_inner.go
   @Author: khaosles
   @Time: 2023/8/10 00:04
   @Desc:
*/

func errInvalidStruct(v any) error {
	return fmt.Errorf("invalid struct %v", v)
}
