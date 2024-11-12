package fileutils

import "fmt"

func FilesFromMap(fileMap map[string]string, extension string) {
	for src, dest := range fileMap {
		content, err := ReadFileContent(src)
		if err != nil {
			fmt.Printf("Error reading file %s : %v\n", src, err)
			continue
		}
		CreateFile(dest, extension, content)
	}
}