package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	code uint32
	err  error
}

func (e *Error) Code() uint32 {
	return e.code
}

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d ; msg = %s", e.code, e.err.Error())
}

func NewErr(errCode uint32, errMsg string) error {
	return &Error{
		code: errCode,
		err:  errors.New(errMsg),
	}
}

func WrapCode(errCode uint32, err error) error {
	return &Error{
		code: errCode,
		err:  errors.Wrap(err, ""),
	}
}
