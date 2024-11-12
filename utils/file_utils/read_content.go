package fileutils

import "os"

func ReadFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	return string(content), err
}