package commands

import (
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
	cmd.Stderr = os.Stderr
	if redirect > 0 {
		output, err := cmd.Output()
		if err != nil {
			return
		}
		writeOutput(string(output), redirect, file)
	} else {
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}
