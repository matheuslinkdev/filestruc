package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(path, ext, content string) {
	filePath := filepath.Join(path)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing on file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Created file: %s\n", filePath)
}


