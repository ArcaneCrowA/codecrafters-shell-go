package commands

import (
	"errors"
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
