package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(path, ext string) {
	filePath := filepath.Join(path)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	fmt.Printf("Created file: %s\n", filePath)
}


