# shell

Spawn a command in the background

## Usage

### Example with `ssh-keygen`

````go
stdOutBuffer, stdErrBuffer, err := shell.Shell(shell.ShellArguments{
    Name:    "ssh-keygen",
    Args:    []string{"-t", "ed25519", "-f", "file", "-C", "your_email@example.org", "-q"},
    Timeout: time.Second * 10,
    Input:   "passcode\npasscode\n",
})
````

### Example with `python`

````go
stdOutBuffer, stdErrBuffer, err := shell.Shell(shell.ShellArguments{
    Name:    "python",
    Args:    []string{"-c", "name = input('What is your name? '); print(f'Hello, {name}!')"},
    Timeout: time.Second * 5,
    Input:   "udfordria\n",
})
````