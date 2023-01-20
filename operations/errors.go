package operations

import "fmt"

type Error struct {
	code int
	err  string
}

var (
	ErrorUnsupportedSchema = Error{
		code: 0,
		err:  "unsupported schema name %v",
	}
	ErrorCantMarshal = Error{
		code: 1,
		err:  "cant marshal %#v, err %v",
	}
	ErrorCantUnMarshal = Error{
		code: 2,
		err:  "cant unmarshal %#v, err %v",
	}
	ErrorUnknownType = Error{
		code: 3,
		err:  "body of envelop is not type %T, has type %T",
	}
	ErrorUnMarshalNotPTR = Error{
		code: 4,
		err:  "please, provide pointer to Interface Element",
	}
	ErrorUnMarshalCantSet = Error{
		code: 5,
		err:  "value for provided interface can not be setted up correctly",
	}
)

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
