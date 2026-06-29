package commands

import (
	"strings"
)

func Echo(args []string, redirect int, file string) {
	writeOutput(strings.Join(args, " "), redirect, file)
}
