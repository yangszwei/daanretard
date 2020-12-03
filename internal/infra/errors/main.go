package errors

import (
	"fmt"
	"strings"
)

// New create an Error
func New(name string, message ...string) Error {
	if len(message) == 0 {
		return Error{
			Name: name,
		}
	}
	return Error{
		Name:    name,
		Message: message[0],
	}
}

// From shorthand for New(err.Error())
func From(err error) Error {
	return New(err.Error())
}

// Is compare the name of two errors, returns true if equal, this does not
// compare the message field since it should not impact how the app handle
// the error
func Is(err1, err2 Error) bool {
	return err1.Name == err2.Name
}

// Error custom error object so errors can be delivered easier, Name is for
// the type of error, for example: "invalid credentials", "unauthorized",
// and the usage of Message is mostly used to specify the scope / fields with
// the error, for example: Error{ Name: "invalid credentials:too long", Message:
// "name" } means the field "name" is too long
type Error struct {
	Name    string
	Message string
}

// SetMessage return a new Error with message
func (e Error) SetMessage(message string) Error {
	e.Message = message
	return e
}

// Error print the error as string (this implement error interface)
func (e Error) Error() string {
	if e.Message == "" {
		return e.Name
	}
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}

// Errors a list of Error
type Errors []Error

// Error prints each error as string, separated by semicolons
func (e Errors) Error() string {
	var errs []string
	for _, err := range e {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, ";")
}
