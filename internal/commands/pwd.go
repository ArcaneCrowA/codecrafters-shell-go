package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func Pwd(redirect int, file string) {
	w, _, cleanup, _ := SetupRedirect(redirect, file)
	defer cleanup()
	fmt.Fprintln(w, pwd())
}

func pwd() string {
	dir, err := filepath.Abs(".")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return dir
}
