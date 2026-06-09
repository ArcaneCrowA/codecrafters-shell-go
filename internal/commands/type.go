package commands

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"slices"
	"strings"
)

var builtin = []string{"exit", "echo", "type"}

func Type(command string) {
	if slices.Contains(builtin, command) {
		fmt.Printf("%s is a shell builtin\n", command)
		return
	}
	if p, e := find(command); e {
		fmt.Printf("%s is %s\n", command, p)
		return
	}

	Invalid(command)
}

func find(command string) (string, bool) {
	paths := strings.SplitSeq(os.Getenv("PATH"), string(os.PathSeparator))
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
		if !file.IsDir() && file.Mode()&0111 != 0 {
			return p, true
		}
	}
	return "", false
}
