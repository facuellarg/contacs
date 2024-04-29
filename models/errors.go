package models

import "fmt"

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewError(message string, code int) *Error {
	return &Error{Message: message, Code: code}
}

func NewInternalError(message string) *Error {
	return NewError(message, 500)
}

func (e Error) Error() string {
	return fmt.Sprintf("Error: %s, Code: %d", e.Message, e.Code)
}
