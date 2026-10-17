package commands

import (
	"fmt"
)

func Invalid(command string, isAppend bool, redirect int, file string) {
	w, _, cleanup, _ := SetupRedirect(isAppend, redirect, file)
	defer cleanup()
	fmt.Fprintf(w, "%s: not found\n", command)
}
