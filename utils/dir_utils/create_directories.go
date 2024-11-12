package dir_utils

import (
	"fmt"
	"os"
)

func CreateDirectories(directories []string) {
	for _, dir := range directories {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating folder %s: %v\n", dir, err)
		} else {
			fmt.Printf("Created folder: %s\n", dir)
		}
	}
}