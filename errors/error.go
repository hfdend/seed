package errors

import (
	"fmt"
)

const (
	System   = 50001
	Business = 40001
	NoLogin  = 1007

	RequestLostTime = 6001
	RequestExpired  = 6002
	Signature       = 6003
)

const (
	RPCNoAuthorization = iota + 8000
)

type Error struct {
	Err  error
	Code int
}

func New(err interface{}, code ...int) error {
	var (
		e  error
		ok bool
		c  int
	)
	if e, ok = err.(error); !ok {
		e = fmt.Errorf("%v", err)
	}
	if len(code) == 0 {
		c = Business
	} else {
		c = code[0]
	}
	return &Error{
		Err:  e,
		Code: c,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.Err.Error())
}
