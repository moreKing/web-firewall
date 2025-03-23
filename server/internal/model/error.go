package model

import "fmt"

// 自定义的错误类型
type apiError struct {
	code int
	msg  string
}

type ApiError = *apiError

func (a apiError) Error() string {
	return fmt.Sprintf("Error: [%d] %s", a.code, a.msg)
}

func (a apiError) Code() int {
	return a.code
}

func (a apiError) Msg() string {
	return a.msg
}

func NewApiError(code int, msg string) ApiError {
	return &apiError{code, msg}
}
