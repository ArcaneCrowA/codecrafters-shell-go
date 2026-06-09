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

		fmt.Printf("%s: command not found", command)
	}
}
