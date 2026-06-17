package commands

import (
	"fmt"
	"slices"
)

var builtin = []string{"exit", "echo", "type", "pwd"}

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
