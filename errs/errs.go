package errs

import (
	"fmt"
)

// ErrInfo matches the error interface but provides more information
type ErrInfo struct {
	// Err is the actual error message
	Err string
	// 0 if the script was successfull, 1 if not
	ExitStatus *int
	// the command that the script was trying to run
	Command string
}

// Error is used for implementing the error interface, and for creating
// a proper error string
func (e *ErrInfo) Error() string {
	if e.ExitStatus != nil {
		return fmt.Sprintf("%s: %s, exit status: %d", e.Err, e.Command, *e.ExitStatus)
	}

	return fmt.Sprintf("%s: command: %s", e.Err, e.Command)
}

// NewOperationFailedErr occurs when an operation attempted to run but failed. It includes the exit status code
func NewOperationFailedErr(status int, command string) *ErrInfo {
	return &ErrInfo{
		Err:        "operation failed",
		ExitStatus: &status,
		Command:    command,
	}
}

// NewCommandFailedErr occurs when an operation failed before it was attempted
func NewCommandFailedErr(command string) *ErrInfo {
	return &ErrInfo{
		Err:     "operation failed",
		Command: command,
	}
}
