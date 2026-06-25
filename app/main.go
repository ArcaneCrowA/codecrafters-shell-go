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
		cleaned := strings.TrimSpace(command)
		args := make([]string, 0, 1)
		var word strings.Builder
		close := true
		for _, r := range cleaned {
			if r == '\'' {
				close = !close
				if close {
					args = append(args, word.String())
					word.Reset()
				}
				continue
			}
			if close && r == ' ' {
				if word.Len() > 0 {
					args = append(args, word.String())
					word.Reset()
				}
				continue
			}
			word.WriteRune(r)
		}
		args = append(args, word.String())
		command = args[0]
		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			commands.Echo(args[1:])
		case "type":
			commands.Type(args[1])
		case "pwd":
			commands.Pwd()
		case "cd":
			commands.Cd(args[1])
		default:
			commands.Exec(args)
		}
	}
}
