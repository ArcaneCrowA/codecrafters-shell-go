package commands

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

func Exec(args []string) {
	_, exists := find(args[0])
	if !exists {
		Invalid(args[0])
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(output[:len(output)-1]))
}
