package operations

import "fmt"

type Error struct {
	code int
	err  string
}

func (e *Error) Error() string {
	return e.err
}

func IsErrorWithCode(e interface{}, code int) bool {
	err, ok := e.(Error)
	if !ok {
		return false
	}
	if err.code == code {
		return true
	}
	return false
}

func (e *Error) formatError(i ...interface{}) *Error {
	return &Error{
		code: e.code,
		err:  fmt.Sprintf(e.err, i),
	}
}
