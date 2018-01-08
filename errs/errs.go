package errs

import (
	"fmt"
)

// ErrInfo matches the error interface but provides more information
type ErrInfo struct {
	// Err is the actual error message
	Err string
	// ErrCode is used for setting the error code
	ErrCode    int
	ExitStatus *int
	Command    string
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
