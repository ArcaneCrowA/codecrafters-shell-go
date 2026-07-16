package commands

import (
	"log/slog"
	"os"
	"os/exec"
)

func Exec(args []string, redirect int, file string) {
	_, exists := find(args[0])
	if !exists {
		Invalid(args[0], redirect, file)
		return
	}

	cmd := exec.Command(args[0], args[1:]...)

	if redirect == 0 {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
		return
	}

	outFile, err := SetupOutputFile(file)
	if err != nil {
		slog.Error("failed to open redirect file", "err", err.Error())
		return
	}
	defer outFile.Close()

	switch redirect {
	case 1:
		cmd.Stdout = outFile
		cmd.Stderr = os.Stderr
	case 2:
		cmd.Stdout = os.Stdout
		cmd.Stderr = outFile
	}

	_ = cmd.Run()
}
