package xerror

import (
	"strings"
)

/*
   @File: api_error.go
   @Author: khaosles
   @Time: 2023/8/13 13:59
   @Desc:
*/

const (
	StatusOK = 20000 // RFC 9110, 15.3.1

	StatusBadRequest       = 40000 // RFC 9110, 15.5.1
	StatusParamError       = 40001 // RFC 9110, 15.5.1
	StatusUnauthorized     = 40100 // RFC 9110, 15.5.2
	StatusNotToken         = 40101 // RFC 9110, 15.5.2
	StatusTokenExpired     = 40102 // RFC 9110, 15.5.2
	StatusTokenIpError     = 40103 // RFC 9110, 15.5.2
	StatusForbidden        = 40300 // RFC 9110, 15.5.4
	StatusNotFound         = 40400 // RFC 9110, 15.5.5
	StatusMethodNotAllowed = 40500 // RFC 9110, 15.5.6

	StatusInternalServerError = 50000 // RFC 9110, 15.6.1
	StatusNotImplemented      = 50100 // RFC 9110, 15.6.2
	StatusBadGateway          = 50200 // RFC 9110, 15.6.3
)

type apiError struct {
	code int
	msg  string
}

func NewApiError(opts ...ApiOption) IError {
	err := apiError{code: StatusInternalServerError, msg: "system error"}
	for _, opt := range opts {
		opt(&err)
	}
	return &err
}

func (err *apiError) Error() string {
	return err.msg
}

func (err *apiError) Code() int {
	return err.code
}

type ApiOption func(*apiError)

func WithCode(code int) ApiOption {
	return func(a *apiError) {
		a.code = code
	}
}

func WithMsg(msg ...string) ApiOption {
	return func(a *apiError) {
		a.msg = strings.Join(msg, ",")
	}
}
