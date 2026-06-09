package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

func main() {
	for {
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		args := strings.Fields(strings.TrimSpace(command))
		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			commands.Echo(args[1:])
		case "type":
			commands.Type(args[1])
		default:
			commands.Exec(args)
		}
	}
}
