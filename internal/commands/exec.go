package commands

import (
	"bytes"
	"os"
	"os/exec"
)

func Exec(args []string, redirect int, file string) {
	_, exists := find(args[0])
	if !exists {
		Invalid(args[0])
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	var stdoutBuf, stderrBuf bytes.Buffer

	switch redirect {
	case 0:
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	case 1:
		cmd.Stdout = &stdoutBuf
		cmd.Stderr = &stderrBuf
		_ = cmd.Run()
		writeOutput(stdoutBuf.String(), redirect, file)
	case 2:
		cmd.Stdout = os.Stdout
		cmd.Stderr = &stderrBuf
		_ = cmd.Run()
		writeOutput(stderrBuf.String(), redirect, file)
	}
}
