package shell

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

// Argument for the `Shell` function
type ShellArguments struct {
	Name    string
	Args    []string
	Timeout time.Duration
	Input   string
	Dir     string
}

// Spawn a command in the background
// If `arg.Timeout` is not defined or is equal to zero then no timeout will be triggered
// If `arg.Input` is not defined or has a length of zero then no input will passed to the interactive command
// If `arg.Dir` is not defined or has a length of zero then the command will be run at the current folder
func Shell(arg ShellArguments) (bytes.Buffer, bytes.Buffer, int, error) {
	cmd := exec.Command(arg.Name, arg.Args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	if len(arg.Dir) != 0 {
		cmd.Dir = arg.Dir
	}

	var stdOutBuffer bytes.Buffer
	var stdErrBuffer bytes.Buffer

	if len(arg.Input) != 0 {
		input, err := cmd.StdinPipe()
		if err != nil {
			return stdOutBuffer, stdErrBuffer, -1, fmt.Errorf("failed to create pipe for STDIN: %s", err)
		}

		fmt.Fprint(input, arg.Input)
	}

	cmd.Stdout, cmd.Stderr = &stdOutBuffer, &stdErrBuffer

	err := cmd.Start()

	if err != nil {
		return stdOutBuffer, stdErrBuffer, -2, fmt.Errorf("failed to start command: %s", err)
	}

	if arg.Timeout != 0 {
		defer time.AfterFunc(arg.Timeout, func() {
			cmd.Process.Kill()
		}).Stop()
	}

	err = cmd.Wait()

	if err != nil {
		exitCode := -3
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		}
		return stdOutBuffer, stdErrBuffer, exitCode, fmt.Errorf("failed to wait command: %s", err)
	}

	return stdOutBuffer, stdErrBuffer, 0, nil
}
