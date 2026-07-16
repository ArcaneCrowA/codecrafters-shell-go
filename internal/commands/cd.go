package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func Cd(path string, redirect int, file string) {
	_, ew, cleanup, _ := SetupRedirect(redirect, file)
	defer cleanup()
	if path == "~" {
		path = os.Getenv("HOME")
	}

	dir, err := filepath.Abs(path)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	if err = os.Chdir(dir); err != nil {
		fmt.Fprintf(ew, "cd: %s: No such file or directory\n", path)
	}
}
