package commands

import (
	"log/slog"
	"os"
	"path/filepath"
)

func Cd(path string) {
	dir, err := filepath.Abs(path)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	os.Chdir(dir)
}
