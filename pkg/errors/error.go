package errors

import (
	"fmt"
)

type Error struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Internal error  `json:"internal"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d msg = %s", e.Code, e.Msg)
}

func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
