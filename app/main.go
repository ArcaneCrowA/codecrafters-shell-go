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
		args := getArgs(command)

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

func getArgs(line string) []string {
	runes := []rune(strings.TrimRight(line, "\r\n"))
	args := make([]string, 0, 1)
	var word strings.Builder
	var split rune
	var open bool
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if open {
			if split == '\'' {
				word.WriteRune(r)
			}
		} else if r == '\\' && i+1 < len(runes) {
			i++
			word.WriteRune(runes[i])
			continue
		}

		if r == '\'' || r == '"' {
			if !open {
				split = r
				open = true
			} else if split == r {
				open = false

			} else {
				word.WriteRune(r)
			}
			continue
		}

		if !open && r == ' ' {
			if word.Len() > 0 {
				args = append(args, word.String())
				word.Reset()
			}
		}
	}

	if word.Len() > 0 {
		args = append(args, word.String())
	}

	return args
}
