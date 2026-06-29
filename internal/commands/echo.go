package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"
)

func Echo(args []string, redirect int, file string) {
	line := strings.Join(args, " ")
	switch redirect {
	case 0:
		fmt.Println(line)
	case 1:
		fullpath := path.Join(pwd(), file)
		if err := os.WriteFile(fullpath, []byte(line), 0644); err != nil {
			slog.Error("failed to write to file", "err", err.Error())
			os.Exit(1)
		}
	}
}
