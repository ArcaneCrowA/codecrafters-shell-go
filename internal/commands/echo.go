package commands

import (
	"strings"
)

func Echo(args []string, file string) {
	writeOutput(strings.Join(args, " "), file)
}
