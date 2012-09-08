package main

import (
	"os"
	"fmt"
	"path/filepath"
	"errors"
)

type project struct {
	WorkingDirectory string
	Name string
}

func newProject() (*project, error) {
	project := new(project)

	workingDirectory, error := os.Getwd()

	if error != nil {
		return project, error
	}

	workingDirectory, error = filepath.Abs(workingDirectory)

	if error != nil {
		return project, error
	}

	name := filepath.Base(workingDirectory)

	project.WorkingDirectory = workingDirectory
	project.Name = name

	return project, errors.New("LOL")
}

func main() {
	project, error := newProject()
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	if project.WorkingDirectory == "/" {
		fmt.Println("Cannot be run from the root!")
		os.Exit(2)
	}

	fmt.Println(project)
}
