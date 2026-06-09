package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		command = strings.TrimSpace(command)
		if command == "exit" {
			os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", command)
	}
}
