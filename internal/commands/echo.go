package commands

import (
	"fmt"
	"strings"
)

func Echo(args []string, redirect int, file string) {
	w, _, cleanup, _ := SetupRedirect(redirect, file)
	defer cleanup()
	fmt.Fprintln(w, strings.Join(args, " "))
}
