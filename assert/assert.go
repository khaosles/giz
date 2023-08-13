package assert

import (
	"strings"
	"time"

	"github.com/khaosles/giz/fileutil"
	"github.com/khaosles/giz/g"
	"github.com/khaosles/giz/xerror"
)

/*
   @File: assert.go
   @Author: khaosles
   @Time: 2023/8/13 13:55
   @Desc:
*/

func IsNull(obj any, code int, msg string) {
	if obj == nil {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsNotNull(obj any, code int, msg string) {
	if obj != nil {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsBlank(s string, code int, msg string) {
	if strings.TrimSpace(s) == "" {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsNotBlank(s string, code int, msg string) {
	if strings.TrimSpace(s) != "" {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsZero[T g.Numeric](num T, code int, msg string) {
	if num == 0 {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsTrue(expr bool, code int, msg string) {
	if expr == true {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsFalse(expr bool, code int, msg string) {
	if expr == false {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsEmpty[T any](arr []T, code int, msg string) {
	if len(arr) == 0 {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsEmptyPointer(obj any, code int, msg string) {
	if obj == nil {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsNotEmptyPointer(obj any, code int, msg string) {
	if obj != nil {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsEmptyTime(t time.Time, code int, msg string) {
	if t.IsZero() {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsFile(file string, code int, msg string) {
	if fileutil.IsFile(file) {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsNotFile(file string, code int, msg string) {
	if !fileutil.IsFile(file) {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func IsNotImplemented(point any, code int, msg string) {
	if point == nil {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func ExecSuccess(err error, code int, msg string) {
	if err != nil {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}

func CheckCount(count int, code int, msg string) {
	if count < 1 {
		panic(xerror.NewApiError(xerror.WithCode(code), xerror.WithMsg(msg)))
	}
}
