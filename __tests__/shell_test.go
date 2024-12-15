package __tests__

import (
	"testing"
	"time"

	"github.com/udfordria/shell"
)

func TestSSHKeygenCommand(t *testing.T) {
	stdOutBuffer, stdErrBuffer, err := shell.Shell(shell.ShellArguments{
		Name:    "ssh-keygen",
		Args:    []string{"-t", "ed25519", "-f", "file", "-C", "your_email@example.org", "-q"},
		Timeout: time.Second * 10,
		Input:   "passcode\npasscode\n",
	})

	t.Log(stdOutBuffer.String())
	t.Log(stdErrBuffer.String())

	if err != nil {
		panic(err)
	}

}

func TestPythonCommand(t *testing.T) {
	stdOutBuffer, stdErrBuffer, err := shell.Shell(shell.ShellArguments{
		Name:    "python",
		Args:    []string{"-c", "name = input('What is your name? '); print(f'Hello, {name}!')"},
		Timeout: time.Second * 5,
		Input:   "udfordria\n",
	})

	t.Log(stdOutBuffer.String())
	t.Log(stdErrBuffer.String())

	if err != nil {
		panic(err)
	}
}
