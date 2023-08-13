package xerror

/*
   @File: ierror.go
   @Author: khaosles
   @Time: 2023/8/13 13:57
   @Desc:
*/

type IError interface {
	Error() string
	Code() int
}
