package commands

import (
	"fmt"
	"slices"
)

var builtin = []string{"exit", "echo", "type", "pwd"}

func Type(command string, redirect int, file string) {
	w, _, cleanup, _ := SetupRedirect(redirect, file)
	defer cleanup()
	if slices.Contains(builtin, command) {
		fmt.Fprintf(w, "%s is a shell builtin\n", command)
		return
	}
	if p, e := find(command); e {
		fmt.Fprintf(w, "%s is %s\n", command, p)
		return
	}
	fmt.Fprintf(w, "%s: not found\n", command)
}
