package main

import (
	"os"
	"fmt"
	"path/filepath"
	"os/exec"
)

func main() {
	workingDirectory, error := os.Getwd()

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	workingDirectory, error = filepath.Abs(workingDirectory)

	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	os.Setenv("GOPATH", workingDirectory)

	command := exec.Command("go", os.Args[1:]...)
	command.Stdin  = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if error := command.Run(); error != nil {
		fmt.Println(error)
	}
}
