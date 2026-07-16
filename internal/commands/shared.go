package commands

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"
)

func find(command string) (string, bool) {
	paths := strings.SplitSeq(os.Getenv("PATH"), string(os.PathListSeparator))
	for p := range paths {
		fullpath := path.Join(p, command)
		file, exists := exists(fullpath)
		if !exists {
			continue
		}

		if !file.IsDir() && file.Mode()&001 != 0 {
			return fullpath, true
		}
	}
	return "", false
}

func exists(fullpath string) (os.FileInfo, bool) {
	file, err := os.Stat(fullpath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, false
		}
		slog.Error(err.Error())
		os.Exit(1)
	}
	return file, true
}

func writeOutput(line string, redirect int, file string) {
	switch redirect {
	case 0:
		fmt.Println(line)
	case 1:
		var fullpath string
		if path.IsAbs(file) {
			fullpath = file
		} else {
			fullpath = path.Join(pwd(), file)
		}
		if err := os.WriteFile(fullpath, []byte(line), 0644); err != nil {
			slog.Error("failed to write to file", "err", err.Error())
			os.Exit(1)
		}
	}
}
