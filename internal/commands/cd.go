package commands

import (
	"fmt"
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
	if err = os.Chdir(dir); err != nil {
		fmt.Printf("cd: %s: %s", path, err)
	}
}
