package commands

import (
	"fmt"
)

func Invalid(command string, redirect int, file string) {
	w, _, cleanup, _ := SetupRedirect(redirect, file)
	defer cleanup()
	fmt.Fprintf(w, "%s: not found\n", command)
}
