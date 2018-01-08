package errs

import (
	"fmt"
	"net/http"
)

// ErrInfo matches the error interface but provides more information
type ErrInfo struct {
	// Err is the actual error message
	Err string
	// ErrCode is used for setting the error code
	ErrCode    int
	ExitStatus *int
	Command    string
	// StatusCode is used for setting corresponding http response
	// for the error
	StatusCode int `json:"-"`
	// Context is used for giving a detailed information about the error
	context error
}

// Error is used for implementing the error interface, and for creating
// a proper error string
func (e *ErrInfo) Error() string {
	if e.ExitStatus != nil {
		return fmt.Sprintf("%s: %s, exit status: %d", e.Err, e.Command, *e.ExitStatus)
	}

	return fmt.Sprintf("%s: command: %s", e.Err, e.Command)
}

// NewOperationFailedErr ...
func NewOperationFailedErr(status int, command string) *ErrInfo {
	return &ErrInfo{
		Err:        "operation failed",
		ExitStatus: &status,
		Command:    command,
	}
}

// NewCommandFailedErr ...
func NewCommandFailedErr(command string) *ErrInfo {
	return &ErrInfo{
		Err:     "operation failed",
		Command: command,
	}
}

// ErrorWithContext is used for creating errors with their context. For instance, we can have
// an ErrSendGrid error, and we can define specific error reason into ErrSendGrid as a context.
// So it becomes: send grid error: could not connect to api
func ErrorWithContext(parentErr, err error) *ErrInfo {
	castErr, ok := parentErr.(*ErrInfo)
	customErr := new(ErrInfo)

	if ok {
		*customErr = *castErr
	} else {
		customErr.Err = parentErr.Error()
	}

	customErr.context = err

	return customErr
}

// ErrorWithStatus is used for creating errors with their corresponding http status codes.
func ErrorWithStatus(message string, errorCode int, statusCode int) *ErrInfo {
	return &ErrInfo{
		Err:        message,
		ErrCode:    errorCode,
		StatusCode: statusCode,
	}
}

// Context is used for creating a new instance of the error and appending a context into that
func (e ErrInfo) Context(err error) *ErrInfo {
	ctxErr := new(ErrInfo)
	*ctxErr = e
	ctxErr.context = err

	return ctxErr
}

var errOperationFailed = &ErrInfo{Err: "operation failed"}
var errCommandFailed = &ErrInfo{Err: "command failed"}

// ErrMarshalData is sent when marshal could not work
var ErrMarshalData = ErrorWithStatus("could not marshal data", 102, http.StatusBadRequest)

// ErrServiceRequest fires when a request is unable to reach its target
var ErrServiceRequest = ErrorWithStatus("service request error", 105, http.StatusInternalServerError)

// ErrDecoder defines a common decoder error
var ErrDecoder = ErrorWithStatus("could not decode data", 101, http.StatusBadRequest)
