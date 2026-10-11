package commands

import (
	"errors"
	"io"
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

func SetupOutputFile(filename string) (*os.File, error) {
	fullpath := filename
	if !path.IsAbs(fullpath) {
		pwd, _ := os.Getwd()
		fullpath = path.Join(pwd, fullpath)
	}
	if err := os.MkdirAll(path.Dir(fullpath), 0755); err != nil {
		return nil, err
	}
	return os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
}

func SetupAppendFile(filename string) (*os.File, error) {
	fullpath := filename
	if !path.IsAbs(fullpath) {
		pwd, _ := os.Getwd()
		fullpath = path.Join(pwd, fullpath)
	}
	return os.OpenFile(fullpath, os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
}

func SetupRedirect(isAppend bool, redirect int, filename string) (stdout io.Writer, stderr io.Writer, cleanup func(), err error) {
	if redirect == 0 {
		return os.Stdout, os.Stderr, func() {}, nil
	}

	var f *os.File

	if isAppend {
		f, err = SetupAppendFile(filename)
	} else {
		f, err = SetupOutputFile(filename)
	}
	if err != nil {
		return nil, nil, nil, err
	}
	if redirect == 1 {
		return f, os.Stderr, func() { f.Close() }, nil
	}
	return os.Stdout, f, func() { f.Close() }, nil
}
