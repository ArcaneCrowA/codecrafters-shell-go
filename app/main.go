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
	args := make([]string, 0)
	var word strings.Builder
	var split rune
	var open bool

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if r == '\\' {
			if open && split == '"' && i+1 < len(runes) {
				switch runes[i+1] {
				case 't':
					word.WriteRune('\t')
				case 'n':
					word.WriteRune('\n')
				default:
					word.WriteRune(runes[i+1])
				}
				i++
			} else if open && split == '\'' {
				word.WriteRune(r)
			} else if i+1 < len(runes) {
				word.WriteRune(runes[i+1])
				i++
			}
			continue
		}

		if r == '\'' || r == '"' {
			if !open {
				split = r
				open = true
				continue
			} else if split == r {
				open = false
				continue
			}
		}

		if !open && r == ' ' {
			if word.Len() > 0 {
				args = append(args, word.String())
				word.Reset()
			}
			continue
		}
		word.WriteRune(r)
	}

	if word.Len() > 0 {
		args = append(args, word.String())
	}

	return args
}
