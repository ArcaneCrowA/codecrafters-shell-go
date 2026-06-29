package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func Pwd() {
	dir := pwd()
	fmt.Println(dir)
}

func pwd() string {
	dir, err := filepath.Abs(".")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return dir
}
