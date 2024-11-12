package golang

import (
	"fmt"

	"github.com/matheuslinkdev/filestruc/utils/dir_utils"
	fileutils "github.com/matheuslinkdev/filestruc/utils/file_utils"
)

func CreateGoStructure() {
	directories := []string{
		"cmd",
		"cmd/main",
		"pkg",
		"internal",
		"api",
		"configs",
		"scripts",
		"build",
		"docs",
	}

	dir_utils.CreateDirectories(directories)

	fileutils.CreateFile("cmd/main/main.go", "go")

	fmt.Println("Golang folder structure created!")
}