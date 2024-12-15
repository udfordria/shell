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
}

// Spawn a command in the background
// If `arg.Timeout` is not defined or is equal to zero then no timeout will be triggered
// If `arg.Input` is not defined or has a length of zero then no input will passed to the interactive command
func Shell(arg ShellArguments) (bytes.Buffer, bytes.Buffer, error) {
	cmd := exec.Command(arg.Name, arg.Args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	var stdOutBuffer bytes.Buffer
	var stdErrBuffer bytes.Buffer

	if len(arg.Input) != 0 {
		input, err := cmd.StdinPipe()
		if err != nil {
			return stdOutBuffer, stdErrBuffer, fmt.Errorf("failed to create pipe for STDIN: %s", err)
		}

		fmt.Fprint(input, arg.Input)
	}

	cmd.Stdout, cmd.Stderr = &stdOutBuffer, &stdErrBuffer

	err := cmd.Start()

	if err != nil {
		return stdOutBuffer, stdErrBuffer, fmt.Errorf("failed to start command: %s", err)
	}

	if arg.Timeout != 0 {
		defer time.AfterFunc(arg.Timeout, func() {
			cmd.Process.Kill()
		}).Stop()
	}

	err = cmd.Wait()

	if err != nil {
		return stdOutBuffer, stdErrBuffer, fmt.Errorf("failed to wait command: %s", err)
	}

	return stdOutBuffer, stdErrBuffer, nil
}
