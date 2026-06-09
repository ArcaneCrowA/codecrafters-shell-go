package commands

import "fmt"

func Invalid(command string) {
	fmt.Printf("%s: command not found\n", command)
}
