package commands

import (
	"fmt"
	"slices"
)

var builtin = []string{"exit", "echo", "type"}

func Type(command string) {
	if slices.Contains(builtin, command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else {
		Invalid(command)
	}
}
