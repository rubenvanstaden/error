package catch

import (
	"errors"
	"fmt"
	"strings"
)

const (
	CONFLICT = "CONFLICT"
	INTERNAL = "INTERNAL"
	INVALID  = "INVALID"
	NOTFOUND = "NOT_FOUND"
	UNKOWN   = "UNKWON"
)

type Debug struct {
	Op  string
	Err error
}

func (e *Debug) Error() string {
	var b strings.Builder
	if e.Op != "" {
		b.WriteString(string(e.Op))
	}
	if e.Err != nil {
		if b.Len() > 0 {
			b.WriteString(": ")
		}
		b.WriteString(e.Err.Error())
	}
	return b.String()
}

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwraps an application error and returns its code.
func ErrorCode(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Code
	}
	return INTERNAL
}

// Unwraps an application error and returns its message.
func ErrorMessage(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Message
	}
	return "Internal error."
}

// Errorf is a helper function to return an Error with a given code and formatted message.
func Errorf(code string, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
