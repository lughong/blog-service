package errcode

import (
	"fmt"
	"net/http"
)

var codes = map[int]string{}

type Error struct {
	code    int
	msg     string
	details []string
}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Errorf("错误码 %d 已存在，请更换一个", code))
	}

	codes[code] = msg

	return &Error{
		code:    code,
		msg:     msg,
		details: make([]string, 0),
	}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.code, e.msg)
}

func (e *Error) Msgf(args ...interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}

	for _, d := range details {
		e.details = append(e.details, d)
	}

	return e
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case NotFound.Code():
		return http.StatusNotFound
	case UnAuthorization.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	case InternalServerError.Code():
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}
