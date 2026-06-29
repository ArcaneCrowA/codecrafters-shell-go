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
	runes := []rune(line)
	args := make([]string, 0, 1)
	var word strings.Builder
	var split rune
	var open bool
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == '\'' || runes[i] == '"' {
			if !open {
				split = runes[i]
				open = true
			} else if split == runes[i] {
				args = append(args, word.String())
				word.Reset()
				open = false
			} else {
				word.WriteRune(runes[i])
			}
			continue
		}
		if open {
			word.WriteRune(runes[i])
			continue
		}
		switch runes[i] {
		case '\\':
			word.WriteRune(runes[i+1])
			i++
		case ' ':
			args = append(args, word.String())
			word.Reset()
		default:
			word.WriteRune(runes[i])
		}
	}

	if word.Len() > 0 {
		args = append(args, word.String())
	}
	return args
}
