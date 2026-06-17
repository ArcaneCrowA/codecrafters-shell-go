package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func Pwd() {
	dir, err := filepath.Abs(".")
	if err != nil {
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}
	fmt.Println(dir)
}
