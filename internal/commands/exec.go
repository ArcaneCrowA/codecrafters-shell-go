package commands

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Exec(args []string) {
	command, exists := find(args[0])
	if exists {
		cmd := exec.Command(command, args[1:]...)
		output, err := cmd.Output()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		fmt.Printf("$ %s", string(output))
	} else {
		Invalid(command)
	}
}

func find(command string) (string, bool) {
	paths := strings.SplitSeq(os.Getenv("PATH"), string(os.PathListSeparator))
	for p := range paths {
		fullpath := path.Join(p, command)
		file, err := os.Stat(fullpath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			slog.Error(err.Error())
			os.Exit(1)
		}
		if !file.IsDir() && file.Mode()&001 != 0 {
			return fullpath, true
		}
	}
	return "", false
}
