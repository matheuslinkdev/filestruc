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

	fileMap := map[string]string{
		"C:/Users/Matheus/Documents/Projetos/BackEnd/Golang/filestruc/content/golang/main.txt": "cmd/main/main.go",
	}

	fileutils.FilesFromMap(fileMap, "go")

	fmt.Println("Golang folder structure created!")
}