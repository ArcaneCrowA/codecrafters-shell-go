package commands

import (
	"fmt"
	"strings"
)

func Echo(args []string, isAppend bool, redirect int, file string) {
	w, _, cleanup, _ := SetupRedirect(isAppend, redirect, file)
	defer cleanup()
	fmt.Fprintln(w, strings.Join(args, " "))
}
