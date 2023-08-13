package xerror

import (
	"fmt"
	"testing"
)

/*
   @File: api_error_test.go
   @Author: khaosles
   @Time: 2023/8/13 16:26
   @Desc:
*/

func TestNewApiError(t *testing.T) {
	err := NewApiError(WithCode(100), WithMsg("error"))
	e := any(err)
	fmt.Printf("%+v\n", e)
	er, ok := e.(IError)
	if ok {
		fmt.Printf("code=%d, msg=%s\n", er.Code(), er.Error())
	}
}
