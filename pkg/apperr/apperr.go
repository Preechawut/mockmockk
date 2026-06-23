package apperr

import "fmt"

type Error struct {
	Status  int
	Code    string
	Message string
	err     error
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.err)
	}
	return e.Message
}

func (e *Error) Unwrap() error { return e.err }

func Validation(message string) *Error {
	return &Error{Status: 400, Code: "VALIDATION_ERROR", Message: message}
}

func NotFound(code, message string) *Error {
	return &Error{Status: 404, Code: code, Message: message}
}

func Conflict(code, message string) *Error {
	return &Error{Status: 409, Code: code, Message: message}
}

func Internal(err error) *Error {
	return &Error{Status: 500, Code: "INTERNAL_ERROR", Message: "internal error", err: err}
}
