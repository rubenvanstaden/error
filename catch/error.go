package catch

import (
	"bytes"
	"fmt"
)

const (
	ERROR_CONFLICT = "CONFLICT"
	ERROR_INTERNAL = "INTERNAL"
	ERROR_INVALID  = "INVALID"
	ERROR_NOTFOUND = "NOT_FOUND"
)

type Error struct {
	Code    string
	Message string
	Op      string
	Err     error
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	if e.Op != "" {
        fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return ERROR_INTERNAL
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred. Please contact technical support."
}
