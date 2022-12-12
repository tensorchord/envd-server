package server

import (
	"bytes"
	"fmt"
	"net/http"
)

// Error defines a standard application error.
type Error struct {
	// Machine-readable error code.
	HTTPStatusCode int `json:"http_status_code,omitempty"`

	// Human-readable message.
	Message string `json:"message,omitempty"`
	Request string `json:"request,omitempty"`

	// Logical operation and nested error.
	Op  string `json:"op,omitempty"`
	Err error  `json:"error,omitempty"`
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.HTTPStatusCode != 0 {
			fmt.Fprintf(&buf, "<%s> ", http.StatusText(e.HTTPStatusCode))
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

func NewError(code int, err error, op string) error {
	return &Error{
		HTTPStatusCode: code,
		Err:            err,
		Message:        err.Error(),
		Op:             op,
	}
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
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
