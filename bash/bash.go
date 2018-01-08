package bash

import (
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/qhenkart/l2bot/errs"
)

// RunProgram ...
func RunProgram(command string) error {
	args := strings.Split(command, " ")
	return RunCommand(args[0], args[1:]...)
}

// RunCommand ...
func RunCommand(operation string, commands ...string) error {
	cmd := exec.Command(operation, commands...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return errs.NewOperationFailedErr(exiterr.Sys().(syscall.WaitStatus).ExitStatus(), operation+" "+strings.Join(commands, " "))
		}

		return errs.NewCommandFailedErr(operation + " " + strings.Join(commands, " "))

	}
	return nil
}
