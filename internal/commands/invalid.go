package commands

import "fmt"

func Invalid(command string) {
	fmt.Printf("%s: not found\n", command)
}
